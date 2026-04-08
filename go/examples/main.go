package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/manh-cam-studio/dnse-openapi-sdk/go/dnse"
	"github.com/manh-cam-studio/dnse-openapi-sdk/go/websocket"
)

func main() {
	// 1. REST Client Example
	client := dnse.NewClient("your-api-key", "your-api-secret")
	
	status, body, err := client.GetAccounts(true)
	if err != nil {
		log.Fatalf("GetAccounts Error: %v", err)
	}
	fmt.Printf("REST GetAccounts Dry Run - Status: %d, Body Length: %d\n", status, len(body))

	status, body, err = client.GetOHLC("1", map[string]string{"symbol": "SSI"}, true)
	if err != nil {
		log.Fatalf("GetOHLC Error: %v", err)
	}
	fmt.Printf("REST GetOHLC Dry Run - Status: %d, Body Length: %d\n", status, len(body))


	// 2. WebSocket Client Example
	wsClient := websocket.NewTradingClient("your-api-key", "your-api-secret")
	
	wsClient.OnTrade = func(t websocket.Trade) {
		fmt.Printf("Trade event on %s: %f\n", t.Symbol, t.Price)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("Attempting to connect to WebSocket...")
	// This will timeout because we have invalid credentials, but compile-check works.
	err = wsClient.Connect(ctx)
	if err != nil {
		fmt.Printf("WebSocket connect failed (as expected with fake keys): %v\n", err)
	}
}
