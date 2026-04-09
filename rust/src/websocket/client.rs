use std::sync::Arc;
use tokio::sync::{mpsc, Mutex};
use tokio_tungstenite::tungstenite::Message;
use std::error::Error;

use super::auth::create_auth_message;
use super::models::*;
use super::subscriptions::{ChannelConfig, SubscriptionRequest};

use super::transport::{Transport, TungsteniteTransport};
use super::serializer::{get_serializer, Serializer};
use super::dispatcher::DispatcherActor;

#[derive(Clone)]
pub struct Config {
    pub api_key: String,
    pub api_secret: String,
    pub base_url: String,
    pub encoding: String,
}

pub struct TradingClient {
    pub config: Config,
    auth_success: bool,
    pub subscriptions: Arc<Mutex<Vec<ChannelConfig>>>,
    event_tx: Option<mpsc::UnboundedSender<DnseWsEvent>>,
    event_rx: Option<mpsc::UnboundedReceiver<DnseWsEvent>>,
    outbound_tx: Option<mpsc::UnboundedSender<Message>>,
}

impl TradingClient {
    pub fn new(api_key: &str, api_secret: &str) -> Self {
        let (ev_tx, ev_rx) = mpsc::unbounded_channel();
        Self {
            config: Config {
                api_key: api_key.to_string(),
                api_secret: api_secret.to_string(),
                base_url: "wss://ws-openapi.dnse.com.vn".to_string(),
                encoding: "json".to_string(),
            },
            auth_success: false,
            subscriptions: Arc::new(Mutex::new(Vec::new())),
            event_tx: Some(ev_tx),
            event_rx: Some(ev_rx),
            outbound_tx: None,
        }
    }

    pub fn take_receiver(&mut self) -> Option<mpsc::UnboundedReceiver<DnseWsEvent>> {
        self.event_rx.take()
    }

    pub async fn send_raw(&self, req: SubscriptionRequest) -> Result<(), Box<dyn Error>> {
        let msg = serde_json::to_string(&req)?;
        if let Some(tx) = &self.outbound_tx {
            let _ = tx.send(Message::Text(msg.into()));
        }
        Ok(())
    }

    pub async fn connect(&mut self) -> Result<(), Box<dyn Error>> {
        let url_str = format!("{}/v1/stream?encoding={}", self.config.base_url, self.config.encoding);

        let mut transport = Box::new(TungsteniteTransport::new());
        transport.connect(&url_str).await?;

        // Extract reader and writer pipes 
        // We use a local channel for outbound routing to allow multiple actors to write to transport
        let (out_tx, mut out_rx) = mpsc::unbounded_channel::<Message>();
        self.outbound_tx = Some(out_tx.clone());

        let mut in_rx = transport.take_receiver().ok_or("Failed to acquire input receiver")?;

        // Authenticate (blocking)
        let auth_msg = create_auth_message(&self.config.api_key, &self.config.api_secret);
        transport.send(Message::Text(serde_json::to_string(&auth_msg)?.into())).await?;

        // Read the welcome and auth messages
        for _ in 0..2 {
            if let Some(msg) = in_rx.recv().await {
                if let Ok(text) = msg.to_text() {
                    let raw: serde_json::Value = serde_json::from_str(text).unwrap_or(serde_json::json!({}));
                    if let Some(action) = raw.get("action").or_else(|| raw.get("a")).and_then(|v| v.as_str()) {
                        if action == "auth_success" {
                            self.auth_success = true;
                        } else if action == "auth_error" || action == "error" {
                            return Err(Box::from(format!("Auth error: {}", text)));
                        }
                    }
                }
            }
        }

        if !self.auth_success {
            return Err(Box::from("Failed to authenticate (no auth_success received)"));
        }

        // Subscriptions restore
        let subs = self.subscriptions.lock().await;
        if !subs.is_empty() {
            let req = SubscriptionRequest {
                action: "subscribe".to_string(),
                channels: subs.clone(),
            };
            let _ = out_tx.send(Message::Text(serde_json::to_string(&req)?.into()));
        }

        // Dedicated transport writer loop
        tokio::spawn(async move {
            while let Some(msg) = out_rx.recv().await {
                if transport.send(msg).await.is_err() {
                    break;
                }
            }
        });

        // Heartbeat worker
        let out_tx_hb = out_tx.clone();
        tokio::spawn(async move {
            let mut interval = tokio::time::interval(std::time::Duration::from_secs(25));
            loop {
                interval.tick().await;
                if out_tx_hb.send(Message::Text(serde_json::json!({"action": "ping"}).to_string().into())).is_err() {
                    break;
                }
            }
        });

        // Construct Solid Components
        let serializer: Arc<dyn Serializer> = get_serializer(&self.config.encoding).into();
        let dispatcher = DispatcherActor::new(
            serializer,
            self.event_tx.clone().unwrap(),
            out_tx.clone()
        );

        // Start dispatcher event loop non-blocking
        dispatcher.start(in_rx);

        Ok(())
    }
}
