use futures_util::{SinkExt, StreamExt};
use std::sync::Arc;
use tokio::sync::{mpsc, Mutex};
use tokio_tungstenite::{connect_async, tungstenite::Message};
use url::Url;

use super::auth::create_auth_message;
use super::models::*;
use super::subscriptions::{ChannelConfig, SubscriptionRequest};

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

    pub async fn send_raw(&self, req: SubscriptionRequest) -> Result<(), Box<dyn std::error::Error>> {
        let msg = serde_json::to_string(&req)?;
        if let Some(tx) = &self.outbound_tx {
            let _ = tx.send(Message::Text(msg.into()));
        }
        Ok(())
    }

    pub async fn connect(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        let url_str = format!("{}/v1/stream?encoding={}", self.config.base_url, self.config.encoding);
        let url = Url::parse(&url_str)?;

        let (ws_stream, _) = connect_async(url).await?;
        let (mut write, mut read) = ws_stream.split();

        if let Some(msg) = read.next().await {
            let msg = msg?;
            let _ = msg.to_text()?;
        }

        let auth_msg = create_auth_message(&self.config.api_key, &self.config.api_secret);
        write.send(Message::Text(serde_json::to_string(&auth_msg)?.into())).await?;

        if let Some(msg) = read.next().await {
            let msg = msg?;
            let text = msg.to_text()?;
            let raw: serde_json::Value = serde_json::from_str(text)?;
            
            let action = raw.get("action").or_else(|| raw.get("a")).and_then(|v| v.as_str());
            if action == Some("auth_success") {
                self.auth_success = true;
            } else {
                return Err(Box::from(format!("Auth failed: {:?}", text)));
            }
        }

        // Outbound queue setup
        let (out_tx, mut out_rx) = mpsc::unbounded_channel::<Message>();
        self.outbound_tx = Some(out_tx.clone());

        let write_arc = Arc::new(Mutex::new(write));
        let w_clone1 = Arc::clone(&write_arc);

        // Subscriptions
        let subs = self.subscriptions.lock().await;
        if !subs.is_empty() {
            let req = SubscriptionRequest {
                action: "subscribe".to_string(),
                channels: subs.clone(),
            };
            let _ = out_tx.send(Message::Text(serde_json::to_string(&req)?.into()));
        }

        // Outbound sender loop
        tokio::spawn(async move {
            while let Some(msg) = out_rx.recv().await {
                let mut w = w_clone1.lock().await;
                if w.send(msg).await.is_err() { break; }
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

        // Inbox event loop
        let tx = self.event_tx.clone().unwrap();
        let encoding = self.config.encoding.clone();
        
        tokio::spawn(async move {
            while let Some(msg) = read.next().await {
                let msg = match msg {
                    Ok(m) => m,
                    Err(_) => break,
                };
                if msg.is_close() { break; }
                if msg.is_ping() || msg.is_pong() { continue; }

                if let Ok(text) = msg.to_text() {
                    if encoding == "json" {
                        if let Ok(raw) = serde_json::from_str::<serde_json::Value>(text) {
                            let action = raw.get("action").or_else(|| raw.get("a")).and_then(|v| v.as_str());
                            if action == Some("ping") {
                                let _ = out_tx.send(Message::Text(serde_json::json!({"action": "pong"}).to_string().into()));
                                continue;
                            }
                            if action == Some("pong") || action == Some("subscribed") { continue; }

                            let t = raw.get("T").and_then(|v| v.as_str()).unwrap_or("");
                            let event: Option<DnseWsEvent> = match t {
                                "t" => serde_json::from_str(text).map(DnseWsEvent::Trade).ok(),
                                "te" => serde_json::from_str(text).map(DnseWsEvent::TradeExtra).ok(),
                                "q" => serde_json::from_str(text).map(DnseWsEvent::Quote).ok(),
                                "b" => serde_json::from_str(text).map(DnseWsEvent::Ohlc).ok(),
                                "bc" => serde_json::from_str(text).map(DnseWsEvent::OhlcClosed).ok(),
                                "sd" => serde_json::from_str(text).map(DnseWsEvent::SecurityDefinition).ok(),
                                "e" => serde_json::from_str(text).map(DnseWsEvent::ExpectedPrice).ok(),
                                "o" => serde_json::from_str(text).map(DnseWsEvent::Order).ok(),
                                "p" => serde_json::from_str(text).map(DnseWsEvent::Position).ok(),
                                "f" => serde_json::from_str(text).map(DnseWsEvent::ForeignInvestor).ok(),
                                "a" => serde_json::from_str(text).map(DnseWsEvent::AccountUpdate).ok(),
                                "mi" => serde_json::from_str(text).map(DnseWsEvent::MarketIndex).ok(),
                                _ => None,
                            };
                            if let Some(e) = event { let _ = tx.send(e); }
                        }
                    }
                }
            }
        });

        Ok(())
    }
}
