use super::client::DnseClient;
use reqwest::Method;
use std::collections::HashMap;

impl DnseClient {
    pub async fn get_accounts(&self, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        self.request(Method::GET, "/accounts", None, None, None, dry_run).await
    }

    pub async fn get_balances(&self, account_no: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        self.request(Method::GET, &format!("/accounts/{}/balances", account_no), None, None, None, dry_run).await
    }

    pub async fn get_loan_packages(&self, account_no: &str, market_type: &str, symbol: Option<&str>, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        if let Some(s) = symbol {
            query.insert("symbol", s.to_string());
        }
        self.request(Method::GET, &format!("/accounts/{}/loan-packages", account_no), Some(&query), None, None, dry_run).await
    }

    pub async fn get_positions(&self, account_no: &str, market_type: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        self.request(Method::GET, &format!("/accounts/{}/positions", account_no), Some(&query), None, None, dry_run).await
    }

    pub async fn get_position_by_id(&self, market_type: &str, position_id: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        self.request(Method::GET, &format!("/accounts/positions/{}", position_id), Some(&query), None, None, dry_run).await
    }

    pub async fn get_ppse(&self, account_no: &str, market_type: &str, symbol: &str, price: f64, loan_package_id: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());
        query.insert("symbol", symbol.to_string());
        query.insert("price", price.to_string());
        query.insert("loanPackageId", loan_package_id.to_string());
        self.request(Method::GET, &format!("/accounts/{}/ppse", account_no), Some(&query), None, None, dry_run).await
    }

    pub async fn close_position(&self, position_id: &str, market_type: &str, trading_token: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let mut query = HashMap::new();
        query.insert("marketType", market_type.to_string());

        let mut headers = HashMap::new();
        headers.insert("trading-token", trading_token.to_string());

        self.request(Method::POST, &format!("/accounts/positions/{}/close", position_id), Some(&query), None, Some(&headers), dry_run).await
    }
}
