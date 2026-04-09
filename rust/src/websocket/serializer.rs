use crate::websocket::models::DnseWsEvent;
use async_trait::async_trait;

#[async_trait]
pub trait Serializer: Send + Sync {
    fn parse_event(&self, text: &str) -> Option<DnseWsEvent>;
    fn extract_action(&self, text: &str) -> Option<String>;
    fn name(&self) -> &str;
}

pub struct JsonSerializer;

impl JsonSerializer {
    pub fn new() -> Self {
        Self
    }
}

#[async_trait]
impl Serializer for JsonSerializer {
    fn extract_action(&self, text: &str) -> Option<String> {
        let raw: Result<serde_json::Value, _> = serde_json::from_str(text);
        if let Ok(value) = raw {
            value.get("action").or_else(|| value.get("a")).and_then(|v| v.as_str()).map(|s| s.to_string())
        } else {
            None
        }
    }

    fn parse_event(&self, text: &str) -> Option<DnseWsEvent> {
        let raw: Result<serde_json::Value, _> = serde_json::from_str(text);
        if let Ok(value) = raw {
            let t = value.get("T").and_then(|v| v.as_str()).unwrap_or("");
            match t {
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
            }
        } else {
            None
        }
    }

    fn name(&self) -> &str {
        "json"
    }
}

pub fn get_serializer(encoding: &str) -> Box<dyn Serializer> {
    // Currently only json is implemented. Fallback to json.
    if encoding == "msgpack" {
        // We can implement MsgpackSerializer later
        Box::new(JsonSerializer::new())
    } else {
        Box::new(JsonSerializer::new())
    }
}
