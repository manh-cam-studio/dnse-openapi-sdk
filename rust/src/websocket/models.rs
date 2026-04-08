#![allow(non_snake_case)]

use serde::{de, Deserialize, Deserializer, Serialize};
use serde_json::Value;

pub fn deserialize_flex_f64<'de, D>(deserializer: D) -> Result<f64, D::Error>
where
    D: Deserializer<'de>,
{
    let value: Value = Deserialize::deserialize(deserializer)?;
    match value {
        Value::Number(n) => n.as_f64().ok_or_else(|| de::Error::custom("Invalid f64")),
        Value::String(s) => s.parse::<f64>().map_err(de::Error::custom),
        _ => Ok(0.0),
    }
}

pub fn deserialize_flex_i32<'de, D>(deserializer: D) -> Result<i32, D::Error>
where
    D: Deserializer<'de>,
{
    let value: Value = Deserialize::deserialize(deserializer)?;
    match value {
        Value::Number(n) => n.as_i64().map(|v| v as i32).ok_or_else(|| de::Error::custom("Invalid i32")),
        Value::String(s) => s.parse::<i32>().map_err(de::Error::custom),
        _ => Ok(0),
    }
}

pub fn deserialize_flex_i64<'de, D>(deserializer: D) -> Result<i64, D::Error>
where
    D: Deserializer<'de>,
{
    let value: Value = Deserialize::deserialize(deserializer)?;
    match value {
        Value::Number(n) => n.as_i64().ok_or_else(|| de::Error::custom("Invalid i64")),
        Value::String(s) => s.parse::<i64>().map_err(de::Error::custom),
        _ => Ok(0),
    }
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub enum DnseWsEvent {
    Trade(Trade),
    TradeExtra(TradeExtra),
    ExpectedPrice(ExpectedPrice),
    SecurityDefinition(SecurityDefinition),
    Quote(Quote),
    Ohlc(Ohlc),
    OhlcClosed(OhlcClosed),
    Order(Order),
    Position(Position),
    AccountUpdate(AccountUpdate),
    MarketIndex(MarketIndex),
    ForeignInvestor(ForeignInvestor),
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct PriceLevel {
    #[serde(alias = "price", default, deserialize_with = "deserialize_flex_f64")]
    pub price: f64,
    #[serde(alias = "qtty", default, deserialize_with = "deserialize_flex_i64")]
    pub quantity: i64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct Trade {
    pub marketId: String,
    pub boardId: String,
    pub isin: String,
    pub symbol: String,
    #[serde(rename = "matchPrice", default, deserialize_with = "deserialize_flex_f64")]
    pub price: f64,
    #[serde(rename = "matchQtty", default, deserialize_with = "deserialize_flex_i64")]
    pub quantity: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalVolumeTraded: i64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub grossTradeAmount: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub highestPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub lowestPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub openPrice: f64,
    #[serde(default)]
    pub tradingSessionId: String,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct TradeExtra {
    pub marketId: String,
    pub boardId: String,
    pub isin: String,
    pub symbol: String,
    #[serde(rename = "matchPrice", default, deserialize_with = "deserialize_flex_f64")]
    pub price: f64,
    #[serde(rename = "matchQtty", default, deserialize_with = "deserialize_flex_i64")]
    pub quantity: i64,
    #[serde(default)]
    pub side: String,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub avgPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalVolumeTraded: i64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub grossTradeAmount: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub highestPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub lowestPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub openPrice: f64,
    #[serde(default)]
    pub tradingSessionId: String,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct ForeignInvestor {
    pub marketId: String,
    pub boardId: String,
    pub symbol: String,
    #[serde(default)]
    pub tradingSessionId: String,
    #[serde(default)]
    pub transactTime: String,
    #[serde(default)]
    pub foreignInvestorTypeCode: String,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub sellVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub sellTradedAmount: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub buyVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub buyTradedAmount: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalSellVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalSellTradedAmount: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalBuyVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalBuyTradedAmount: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub foreignerOrderLimitQuantity: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub foreignerBuyPossibleQuantity: i64,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct MarketIndex {
    pub indexName: String,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub changedRatio: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub changedValue: f64,
    
    #[serde(default, deserialize_with = "deserialize_flex_i32")]
    pub fluctuationSteadinessIssueCount: i32,
    #[serde(default, deserialize_with = "deserialize_flex_i32")]
    pub fluctuationDownIssueCount: i32,
    #[serde(default, deserialize_with = "deserialize_flex_i32")]
    pub fluctuationUpIssueCount: i32,
    #[serde(default, deserialize_with = "deserialize_flex_i32")]
    pub fluctuationLowerLimitIssueCount: i32,
    #[serde(default, deserialize_with = "deserialize_flex_i32")]
    pub fluctuationUpperLimitIssueCount: i32,
    
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub fluctuationDownIssueVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub fluctuationUpIssueVolume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub fluctuationSteadinessIssueVolume: i64,

    #[serde(default)]
    pub currencyCode: String,
    #[serde(default)]
    pub indexTypeCode: String,

    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub lowestValueIndexes: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub highestValueIndexes: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub priorValueIndexes: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub valueIndexes: f64,
    
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub contauctAccTrdVal: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub contauctAccTrdVol: i64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub blkTrdAccTrdVal: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub blkTrdAccTrdVol: i64,

    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub grossTradeAmount: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub totalVolumeTraded: i64,

    #[serde(default)]
    pub marketId: String,
    #[serde(default)]
    pub tradingSessionId: String,
    pub transactTime: Option<String>,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct ExpectedPrice {
    pub marketId: String,
    pub boardId: String,
    pub symbol: String,
    pub isin: Option<String>,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub closePrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub expectedTradePrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub expectedTradeQuantity: i64,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct SecurityDefinition {
    pub marketId: String,
    pub boardId: String,
    pub symbol: String,
    pub isin: Option<String>,
    pub productGrpId: Option<String>,
    pub securityGroupId: Option<String>,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub basicPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub ceilingPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub floorPrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub openInterestQuantity: i64,
    pub securityStatus: Option<String>,
    pub symbolAdminStatusCode: Option<String>,
    pub symbolTradingMethodStatusCode: Option<String>,
    pub symbolTradingSanctionStatusCode: Option<String>,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct Quote {
    pub marketId: Option<String>,
    pub boardId: String,
    pub symbol: String,
    pub isin: Option<String>,
    #[serde(default)]
    pub bid: Vec<PriceLevel>,
    #[serde(default)]
    pub offer: Vec<PriceLevel>,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub totalOfferQtty: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub totalBidQtty: f64,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct Ohlc {
    pub symbol: String,
    pub resolution: String,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub open: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub high: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub low: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub close: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub volume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub time: i64,
    #[serde(rename = "type")]
    pub kind: String,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub lastUpdated: i64,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct OhlcClosed {
    pub symbol: String,
    pub resolution: String,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub open: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub high: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub low: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub close: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub volume: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub time: i64,
    #[serde(rename = "type")]
    pub kind: String,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub lastUpdated: i64,
    #[serde(default)]
    pub _receivedAt: f64,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct Order {
    pub id: String,
    pub side: String,
    pub accountNo: String,
    pub symbol: String,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub price: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub priceSecure: f64,
    #[serde(default, deserialize_with = "deserialize_flex_f64")]
    pub averagePrice: f64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub quantity: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub fillQuantity: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub canceledQuantity: i64,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub leaveQuantity: i64,
    #[serde(default)]
    pub orderType: String,
    #[serde(default)]
    pub orderStatus: String,
    #[serde(default)]
    pub loanPackageId: Option<i64>,
    #[serde(default)]
    pub marketType: Option<String>,
    #[serde(default)]
    pub transDate: Option<String>,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct Position {
    pub symbol: String,
    #[serde(default, deserialize_with = "deserialize_flex_i64")]
    pub quantity: i64,
    pub averagePrice: String,
    pub marketValue: Option<String>,
    pub costBasis: Option<String>,
    pub unrealizedPl: Option<String>,
    pub unrealizedPlPercent: Option<String>,
}

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct AccountUpdate {
    pub cash: String,
    pub buyingPower: String,
    pub portfolioValue: Option<String>,
    pub equity: String,
}
