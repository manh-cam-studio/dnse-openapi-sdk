use super::client::DnseClient;
use reqwest::Method;
use std::collections::HashMap;

impl DnseClient {
    pub async fn get_security_definition(&self, symbol: &str, board_id: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        if let Some(b) = board_id {
            query.insert("boardId", b.to_string());
        }
        self.request(Method::GET, &format!("/price/{}/secdef", symbol), Some(&query), None, None, dry_run).await
    }

    pub async fn get_ohlc(&self, bar_type: &str, mut query_params: HashMap<&str, String>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        query_params.insert("type", bar_type.to_string());
        self.request(Method::GET, "/price/ohlc", Some(&query_params), None, None, dry_run).await
    }

    pub async fn get_trades(&self, symbol: &str, board_id: Option<&str>, from_date: Option<&str>, to_date: Option<&str>, limit: Option<usize>, order: Option<&str>, next_page_token: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        if let Some(b) = board_id { query.insert("boardId", b.to_string()); }
        if let Some(f) = from_date { query.insert("from", f.to_string()); }
        if let Some(t) = to_date { query.insert("to", t.to_string()); }
        if let Some(l) = limit { query.insert("limit", l.to_string()); }
        if let Some(o) = order { query.insert("order", o.to_string()); }
        if let Some(n) = next_page_token { query.insert("nextPageToken", n.to_string()); }

        self.request(Method::GET, &format!("/price/{}/trades", symbol), Some(&query), None, None, dry_run).await
    }

    pub async fn get_instruments(&self, symbol: Option<&str>, market_id: Option<&str>, security_group_id: Option<&str>, index_name: Option<&str>, limit: Option<usize>, page: Option<usize>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        if let Some(s) = symbol { query.insert("symbol", s.to_string()); }
        if let Some(m) = market_id { query.insert("marketId", m.to_string()); }
        if let Some(g) = security_group_id { query.insert("securityGroupId", g.to_string()); }
        if let Some(i) = index_name { query.insert("indexName", i.to_string()); }
        if let Some(l) = limit { query.insert("limit", l.to_string()); }
        if let Some(p) = page { query.insert("page", p.to_string()); }

        self.request(Method::GET, "/instruments", Some(&query), None, None, dry_run).await
    }

    pub async fn get_latest_trade(&self, symbol: &str, board_id: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        if let Some(b) = board_id {
            query.insert("boardId", b.to_string());
        }
        self.request(Method::GET, &format!("/price/{}/trades/latest", symbol), Some(&query), None, None, dry_run).await
    }

    pub async fn get_close_price(&self, symbol: &str, board_id: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        if let Some(b) = board_id {
            query.insert("boardId", b.to_string());
        }
        self.request(Method::GET, &format!("/price/{}/close", symbol), Some(&query), None, None, dry_run).await
    }
}
