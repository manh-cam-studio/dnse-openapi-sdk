use std::sync::Arc;
use tokio::sync::mpsc;
use tokio_tungstenite::tungstenite::Message;

use crate::websocket::models::DnseWsEvent;
use crate::websocket::serializer::Serializer;

pub struct DispatcherActor {
    serializer: Arc<dyn Serializer>,
    event_tx: mpsc::UnboundedSender<DnseWsEvent>,
    outbound_tx: mpsc::UnboundedSender<Message>,
}

impl DispatcherActor {
    pub fn new(
        serializer: Arc<dyn Serializer>,
        event_tx: mpsc::UnboundedSender<DnseWsEvent>,
        outbound_tx: mpsc::UnboundedSender<Message>,
    ) -> Self {
        Self {
            serializer,
            event_tx,
            outbound_tx,
        }
    }

    pub fn start(self, mut in_rx: mpsc::UnboundedReceiver<Message>) {
        tokio::spawn(async move {
            while let Some(msg) = in_rx.recv().await {
                if msg.is_close() {
                    break;
                }
                if msg.is_ping() || msg.is_pong() {
                    continue;
                }

                if let Ok(text) = msg.to_text() {
                    let action = self.serializer.extract_action(text);

                    if action.as_deref() == Some("ping") {
                        let _ = self.outbound_tx.send(Message::Text(serde_json::json!({"action": "pong"}).to_string().into()));
                        continue;
                    }
                    if action.as_deref() == Some("pong") || action.as_deref() == Some("subscribed") {
                        continue;
                    }

                    if let Some(event) = self.serializer.parse_event(text) {
                        let _ = self.event_tx.send(event);
                    }
                }
            }
        });
    }
}
