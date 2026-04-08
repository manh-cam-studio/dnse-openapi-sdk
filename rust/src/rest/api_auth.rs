use super::client::DnseClient;
use reqwest::Method;
use serde_json::json;

impl DnseClient {
    pub async fn create_trading_token(&self, otp_type: &str, passcode: &str, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let body = json!({
            "otpType": otp_type,
            "passcode": passcode,
        });
        self.request(Method::POST, "/registration/trading-token", None, Some(&body), None, dry_run).await
    }

    pub async fn send_email_otp(&self, dry_run: bool) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        self.request(Method::POST, "/registration/send-email-otp", None, None, None, dry_run).await
    }
}
