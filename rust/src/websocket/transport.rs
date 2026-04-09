use async_trait::async_trait;
use tokio::sync::mpsc;
use tokio_tungstenite::{connect_async, tungstenite::Message};
use url::Url;
use std::error::Error;
use futures_util::{SinkExt, StreamExt};

#[async_trait]
pub trait Transport: Send + Sync {
    async fn connect(&mut self, url: &str) -> Result<(), Box<dyn Error>>;
    async fn send(&self, msg: Message) -> Result<(), Box<dyn Error>>;
    // Take the receiver channel. It will yield incoming Text messages.
    fn take_receiver(&mut self) -> Option<mpsc::UnboundedReceiver<Message>>;
}

pub struct TungsteniteTransport {
    outbound_tx: Option<mpsc::UnboundedSender<Message>>,
    inbound_rx: Option<mpsc::UnboundedReceiver<Message>>,
}

impl TungsteniteTransport {
    pub fn new() -> Self {
        Self {
            outbound_tx: None,
            inbound_rx: None,
        }
    }
}

#[async_trait]
impl Transport for TungsteniteTransport {
    async fn connect(&mut self, url_str: &str) -> Result<(), Box<dyn Error>> {
        let url = Url::parse(url_str)?;
        let (ws_stream, _) = connect_async(url).await?;
        let (mut write, mut read) = ws_stream.split();

        // One channel for external systems to queue outbound messages
        let (out_tx, mut out_rx) = mpsc::unbounded_channel::<Message>();
        self.outbound_tx = Some(out_tx);

        // One channel to bubble up inbound messages
        let (in_tx, in_rx) = mpsc::unbounded_channel::<Message>();
        self.inbound_rx = Some(in_rx);

        // Outbound sender loop
        tokio::spawn(async move {
            while let Some(msg) = out_rx.recv().await {
                if write.send(msg).await.is_err() {
                    break;
                }
            }
        });

        // Inbound receiver loop
        tokio::spawn(async move {
            while let Some(msg) = read.next().await {
                if let Ok(m) = msg {
                    if in_tx.send(m).is_err() {
                        break;
                    }
                } else {
                    break;
                }
            }
        });

        Ok(())
    }

    async fn send(&self, msg: Message) -> Result<(), Box<dyn Error>> {
        if let Some(tx) = &self.outbound_tx {
            tx.send(msg)?;
            Ok(())
        } else {
            Err(Box::from("Transport not connected or outbound channel missing"))
        }
    }

    fn take_receiver(&mut self) -> Option<mpsc::UnboundedReceiver<Message>> {
        self.inbound_rx.take()
    }
}
