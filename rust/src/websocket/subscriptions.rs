use serde::{Deserialize, Serialize};

use super::client::TradingClient;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ChannelConfig {
    pub name: String,
    pub symbols: Vec<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SubscriptionRequest {
    pub action: String,
    pub channels: Vec<ChannelConfig>,
}

impl TradingClient {
    pub async fn subscribe_channel(&self, channel_name: &str, symbols: Vec<String>) -> Result<(), Box<dyn std::error::Error>> {
        let req = SubscriptionRequest {
            action: "subscribe".to_string(),
            channels: vec![ChannelConfig {
                name: channel_name.to_string(),
                symbols,
            }],
        };

        self.send_raw(req).await
    }

    pub async fn unsubscribe_channel(&self, channel_name: &str, symbols: Vec<String>) -> Result<(), Box<dyn std::error::Error>> {
        let req = SubscriptionRequest {
            action: "unsubscribe".to_string(),
            channels: vec![ChannelConfig {
                name: channel_name.to_string(),
                symbols,
            }],
        };

        self.send_raw(req).await
    }

    pub async fn subscribe_trades(&self, symbols: Vec<String>, board_id: Option<&str>) -> Result<(), Box<dyn std::error::Error>> {
        let boards = match board_id {
            Some(b) => vec![b.to_string()],
            None => vec!["G1", "G3", "G4", "G7", "T1", "T2", "T3", "T4", "T6"].into_iter().map(String::from).collect(),
        };

        for board in boards {
            let channel = format!("tick.{}.{}", board, self.config.encoding);
            self.subscribe_channel(&channel, symbols.clone()).await?;
        }
        Ok(())
    }

    pub async fn subscribe_trade_extra(&self, symbols: Vec<String>, board_id: Option<&str>) -> Result<(), Box<dyn std::error::Error>> {
        let boards = match board_id {
            Some(b) => vec![b.to_string()],
            None => vec!["G1", "G3", "G4", "G7", "T1", "T2", "T3", "T4", "T6"].into_iter().map(String::from).collect(),
        };

        for board in boards {
            let channel = format!("tick_extra.{}.{}", board, self.config.encoding);
            self.subscribe_channel(&channel, symbols.clone()).await?;
        }
        Ok(())
    }


    pub async fn subscribe_ohlc(&self, symbols: Vec<String>, resolution: Option<&str>) -> Result<(), Box<dyn std::error::Error>> {
        let resolutions = match resolution {
            Some(r) => vec![r.to_string()],
            None => vec!["1", "3", "5", "15", "30", "1H", "1D", "1W"].into_iter().map(String::from).collect(),
        };

        for res in resolutions {
            let channel = format!("ohlc.{}.{}", res, self.config.encoding);
            self.subscribe_channel(&channel, symbols.clone()).await?;
        }
        Ok(())
    }

    pub async fn subscribe_ohlc_closed(&self, symbols: Vec<String>, resolution: Option<&str>) -> Result<(), Box<dyn std::error::Error>> {
        let resolutions = match resolution {
            Some(r) => vec![r.to_string()],
            None => vec!["1", "3", "5", "15", "30", "1H", "1D", "1W"].into_iter().map(String::from).collect(),
        };

        for res in resolutions {
            let channel = format!("ohlc_closed.{}.{}", res, self.config.encoding);
            self.subscribe_channel(&channel, symbols.clone()).await?;
        }
        Ok(())
    }

    pub async fn subscribe_quotes(&self, symbols: Vec<String>, board_id: Option<&str>) -> Result<(), Box<dyn std::error::Error>> {
        let boards = match board_id {
            Some(b) => vec![b.to_string()],
            None => vec!["G1", "G2", "G3", "G4", "G5", "G6", "G7"].into_iter().map(String::from).collect(),
        };

        for board in boards {
            let channel = format!("top_price.{}.{}", board, self.config.encoding);
            self.subscribe_channel(&channel, symbols.clone()).await?;
        }
        Ok(())
    }

    pub async fn subscribe_orders(&self) -> Result<(), Box<dyn std::error::Error>> {
        self.subscribe_channel("orders", vec![]).await
    }

    pub async fn subscribe_positions(&self) -> Result<(), Box<dyn std::error::Error>> {
        self.subscribe_channel("positions", vec![]).await
    }

    pub async fn subscribe_account(&self) -> Result<(), Box<dyn std::error::Error>> {
        self.subscribe_channel("account", vec![]).await
    }
}
