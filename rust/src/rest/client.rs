use log::{debug, info};
use reqwest::{Client, Method, RequestBuilder, StatusCode};
use serde_json::Value;
use std::collections::HashMap;
use std::env;
use std::time::Duration;
use uuid::Uuid;
use chrono::Utc;

use super::auth::build_signature;

#[derive(Clone)]
pub struct Config {
    pub api_key: String,
    pub api_secret: String,
    pub base_url: String,
    pub algorithm: String,
    pub hmac_nonce_enabled: bool,
    pub debug: bool,
}

impl Default for Config {
    fn default() -> Self {
        Self {
            api_key: String::new(),
            api_secret: String::new(),
            base_url: "https://openapi.dnse.com.vn".to_string(),
            algorithm: "hmac-sha256".to_string(),
            hmac_nonce_enabled: true,
            debug: false,
        }
    }
}

pub struct DnseClient {
    pub config: Config,
    client: Client,
}

impl DnseClient {
    pub fn new(api_key: String, api_secret: String) -> Self {
        let config = Config {
            api_key,
            api_secret,
            ..Default::default()
        };
        Self::with_config(config)
    }

    pub fn with_config(config: Config) -> Self {
        let client = Client::builder()
            .timeout(Duration::from_secs(60))
            .build()
            .expect("Failed to build HTTP client");
        Self { config, client }
    }

    fn get_date_header(&self) -> String {
        env::var("DATE_HEADER").unwrap_or_else(|_| "Date".to_string())
    }

    fn signature_headers(&self, method: &str, path: &str) -> Result<(String, String), String> {
        let date_value = Utc::now().format("%a, %d %b %Y %H:%M:%S GMT").to_string(); // Format needed for GMT

        let nonce = if self.config.hmac_nonce_enabled {
            Some(Uuid::new_v4().simple().to_string())
        } else {
            None
        };

        let (headers_list, signature) = build_signature(
            &self.config.api_secret,
            method,
            path,
            &date_value,
            &self.config.algorithm,
            nonce.as_deref(),
            &self.get_date_header(),
        )?;

        let mut sig_header_val = format!(
            "Signature keyId=\"{}\",algorithm=\"{}\",headers=\"{}\",signature=\"{}\"",
            self.config.api_key, self.config.algorithm, headers_list, signature
        );

        if let Some(n) = nonce {
            sig_header_val.push_str(&format!(",nonce=\"{}\"", n));
        }

        Ok((date_value, sig_header_val))
    }

    pub async fn request(
        &self,
        method: Method,
        path: &str,
        query: Option<&HashMap<&str, String>>,
        body: Option<&Value>,
        headers: Option<&HashMap<&str, String>>,
        dry_run: bool,
    ) -> Result<(u16, Vec<u8>), Box<dyn std::error::Error>> {
        let debug = self.config.debug || env::var("DEBUG").unwrap_or_default().to_lowercase() == "true";

        let url = format!("{}{}", self.config.base_url, path);
        let mut req_builder = self.client.request(method.clone(), &url);

        if let Some(q) = query {
            req_builder = req_builder.query(q);
        }

        let (date_value, sig_header_value) = self.signature_headers(method.as_str(), path)?;
        let date_header_name = self.get_date_header();

        req_builder = req_builder
            .header(&date_header_name, date_value)
            .header("X-Signature", sig_header_value)
            .header("x-api-key", &self.config.api_key);

        if let Some(b) = body {
            req_builder = req_builder.json(b);
        }

        if let Some(h) = headers {
            for (k, v) in h {
                req_builder = req_builder.header(*k, v);
            }
        }

        let req = req_builder.build()?;

        if dry_run || debug {
            let prefix = if dry_run { "DRY RUN" } else { "DEBUG" };
            info!("{} url: {}", prefix, req.url());
            info!("{} method: {}", prefix, req.method());
            if let Some(q) = query {
                info!("{} query: {:?}", prefix, q);
            }
            info!("{} headers: {:?}", prefix, req.headers());
            if let Some(b) = body {
                info!("{} body: {}", prefix, serde_json::to_string(b)?);
            }
        }

        if dry_run {
            return Ok((200, vec![]));
        }

        let response = self.client.execute(req).await?;
        let status = response.status().as_u16();
        let bytes = response.bytes().await?.to_vec();

        Ok((status, bytes))
    }
}
