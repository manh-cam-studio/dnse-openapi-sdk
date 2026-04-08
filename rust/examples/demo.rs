use dnse_openapi_sdk::rest::DnseClient;
use dnse_openapi_sdk::websocket::TradingClient;
use std::sync::Arc;
use std::time::Duration;

#[tokio::main]
async fn main() {
    // Enable simple logging if needed
    // env_logger::init();

    println!("Starting DNSE OpenAPI Rust SDK Demo...");

    let api_key = "dummy-api-key".to_string();
    let api_secret = "dummy-api-secret".to_string();

    // 1. REST API Client
    let rest_client = DnseClient::new(api_key.clone(), api_secret.clone());
    
    match rest_client.get_accounts(true).await {
        Ok((status, body)) => {
            println!("REST Base `get_accounts` Success: Status={}, BodyLength={}", status, body.len());
        }
        Err(e) => {
            println!("REST Error: {:?}", e);
        }
    }

    // 2. WebSocket Client
    let mut ws_client = TradingClient::new(api_key, api_secret);
    
    ws_client.on_trade = Some(Arc::new(Box::new(|trade| {
        println!("Received Trade Event: {:?} at {}", trade.symbol, trade.price);
    })));

    println!("Attempting to connect to WebSocket...");
    match ws_client.connect().await {
        Ok(_) => {
            println!("Connected successfully to WS!");
            let _ = ws_client.subscribe_trades(vec!["SSI".to_string()], None).await;
            tokio::time::sleep(Duration::from_secs(3)).await;
        }
        Err(e) => {
            println!("WebSocket Connection gracefully failed (expected with spoofed keys): {:?}", e);
        }
    }

    println!("Demo Finished.");
}
