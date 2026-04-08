use base64::{engine::general_purpose::STANDARD as BASE64, Engine};
use chrono::Utc;
use hmac::{Hmac, Mac};
use serde::{Deserialize, Serialize};
use sha2::Sha256;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AuthMessage {
    pub action: String,
    pub api_key: String,
    pub signature: String,
    pub timestamp: i64,
    pub nonce: String,
}

pub fn create_auth_message(api_key: &str, api_secret: &str) -> AuthMessage {
    let now = Utc::now();
    let timestamp = now.timestamp();
    let nonce = now.timestamp_micros().to_string();

    let signature = compute_signature(api_secret, api_key, timestamp, &nonce);

    AuthMessage {
        action: "auth".to_string(),
        api_key: api_key.to_string(),
        signature,
        timestamp,
        nonce,
    }
}

fn compute_signature(api_secret: &str, api_key: &str, timestamp: i64, nonce: &str) -> String {
    let message = format!("{}:{}:{}", api_key, timestamp, nonce);

    let mut mac = Hmac::<Sha256>::new_from_slice(api_secret.as_bytes())
        .expect("HMAC can take key of any size");
    mac.update(message.as_bytes());

    let result = mac.finalize().into_bytes();
    hex::encode(result)
}
