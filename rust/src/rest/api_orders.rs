use super::client::DnseClient;
use reqwest::Method;
use serde_json::Value;
use std::collections::HashMap;

impl DnseClient {
    pub async fn get_orders(&self, account_no: &str, market_type: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(cat) = order_category {
            query.insert("orderCategory", cat.to_string());
        }
        self.request(Method::GET, &format!("/accounts/{}/orders", account_no), Some(&query), None, None, dry_run).await
    }

    pub async fn get_order_detail(&self, account_no: &str, order_id: &str, market_type: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(cat) = order_category {
            query.insert("orderCategory", cat.to_string());
        }
        self.request(Method::GET, &format!("/accounts/{}/orders/{}", account_no, order_id), Some(&query), None, None, dry_run).await
    }

    pub async fn get_execution_detail(&self, account_no: &str, order_id: &str, market_type: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        query.insert("orderCategory", order_category.unwrap_or("NORMAL").to_string());
        
        self.request(Method::GET, &format!("/accounts/{}/executions/{}", account_no, order_id), Some(&query), None, None, dry_run).await
    }

    pub async fn get_order_history(&self, account_no: &str, market_type: &str, from_date: Option<&str>, to_date: Option<&str>, page_size: Option<usize>, page_index: Option<usize>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(f) = from_date { query.insert("from", f.to_string()); }
        if let Some(t) = to_date { query.insert("to", t.to_string()); }
        if let Some(s) = page_size { query.insert("pageSize", s.to_string()); }
        if let Some(i) = page_index { query.insert("pageIndex", i.to_string()); }

        self.request(Method::GET, &format!("/accounts/{}/orders/history", account_no), Some(&query), None, None, dry_run).await
    }

    pub async fn post_order(&self, market_type: &str, payload: &Value, trading_token: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        query.insert("orderCategory", order_category.unwrap_or("NORMAL").to_string());

        let mut headers = HashMap::new();
        headers.insert("trading-token", trading_token.to_string());

        self.request(Method::POST, "/accounts/orders", Some(&query), Some(payload), Some(&headers), dry_run).await
    }

    pub async fn put_order(&self, account_no: &str, order_id: &str, market_type: &str, payload: &Value, trading_token: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(cat) = order_category {
            query.insert("orderCategory", cat.to_string());
        }

        let mut headers = HashMap::new();
        headers.insert("trading-token", trading_token.to_string());

        self.request(Method::PUT, &format!("/accounts/{}/orders/{}", account_no, order_id), Some(&query), Some(payload), Some(&headers), dry_run).await
    }

    pub async fn cancel_order(&self, account_no: &str, order_id: &str, market_type: &str, trading_token: &str, order_category: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(cat) = order_category {
            query.insert("orderCategory", cat.to_string());
        }

        let mut headers = HashMap::new();
        headers.insert("trading-token", trading_token.to_string());

        self.request(Method::DELETE, &format!("/accounts/{}/orders/{}", account_no, order_id), Some(&query), None, Some(&headers), dry_run).await
    }
}
