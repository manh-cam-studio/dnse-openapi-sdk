package websocket

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Config struct {
	APIKey            string
	APISecret         string
	BaseURL           string
	Encoding          string // "json" or "msgpack"
	AutoReconnect     bool
	MaxRetries        int
	HeartbeatInterval time.Duration
	Timeout           time.Duration
}

type TradingClient struct {
	config Config

	transport  Transport
	serializer Serializer
	dispatcher *Dispatcher
	cancelFunc context.CancelFunc

	isAuthenticated bool
	sessionID       string

	done chan struct{}

	subscriptions map[string]SubscriptionRequest
	subsMutex     sync.Mutex

	// Callbacks
	OnTrade              func(Trade)
	OnTradeExtra         func(TradeExtra)
	OnExpectedPrice      func(ExpectedPrice)
	OnSecurityDefinition func(SecurityDefinition)
	OnQuote              func(Quote)
	OnOhlc               func(Ohlc)
	OnOhlcClosed         func(Ohlc)
	OnOrder              func(Order)
	OnPosition           func(Position)
	OnAccountUpdate      func(AccountUpdate)
	OnMarketIndex        func(MarketIndex)
	OnForeignInvestor    func(ForeignInvestor)
	OnError              func(error)
}

func NewTradingClient(apiKey, apiSecret string) *TradingClient {
	return NewTradingClientWithConfig(Config{
		APIKey:            apiKey,
		APISecret:         apiSecret,
		BaseURL:           "wss://ws-openapi.dnse.com.vn",
		Encoding:          "json",
		AutoReconnect:     true,
		MaxRetries:        10,
		HeartbeatInterval: 25 * time.Second,
		Timeout:           60 * time.Second,
	})
}

func NewTradingClientWithConfig(cfg Config) *TradingClient {
	if cfg.BaseURL == "" {
		cfg.BaseURL = "wss://ws-openapi.dnse.com.vn"
	}
	if cfg.Encoding == "" {
		cfg.Encoding = "json"
	}

	client := &TradingClient{
		config:        cfg,
		subscriptions: make(map[string]SubscriptionRequest),
		done:          make(chan struct{}),
		transport:     NewGorillaTransport(),
		serializer:    GetSerializer(cfg.Encoding),
	}

	client.dispatcher = NewDispatcher(client, client.serializer)
	return client
}

func (c *TradingClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	c.cancelFunc = cancel

	url := fmt.Sprintf("%s/v1/stream?encoding=%s", c.config.BaseURL, c.config.Encoding)

	err := c.transport.Connect(ctx, url, c.config.Timeout)
	if err != nil {
		return err
	}

	// Handle initial welcome message
	var welcome map[string]interface{}
	err = c.receiveMessage(&welcome)
	if err != nil {
		return fmt.Errorf("failed to read welcome message: %w", err)
	}

	if sid, ok := welcome["session_id"].(string); ok {
		c.sessionID = sid
	} else if sid, ok := welcome["sid"].(string); ok {
		c.sessionID = sid
	}

	// Authenticate
	err = c.authenticate()
	if err != nil {
		return err
	}

	go c.messageLoop(ctx)

	if c.config.HeartbeatInterval > 0 {
		go c.heartbeatLoop(ctx)
	}

	return nil
}

func (c *TradingClient) authenticate() error {
	authMsg := CreateAuthMessage(c.config.APIKey, c.config.APISecret)

	err := c.sendMessage(authMsg)
	if err != nil {
		return fmt.Errorf("failed to send auth message: %w", err)
	}

	var response map[string]interface{}
	err = c.receiveMessage(&response)
	if err != nil {
		return fmt.Errorf("failed to read auth response: %w", err)
	}

	action, _ := response["action"].(string)
	if action == "" {
		action, _ = response["a"].(string)
	}

	if action == "auth_success" {
		c.isAuthenticated = true
		return nil
	} else if action == "auth_error" || action == "error" {
		msg, _ := response["message"].(string)
		if msg == "" {
			msg, _ = response["msg"].(string)
		}
		return fmt.Errorf("authentication failed: %s", msg)
	}

	return fmt.Errorf("unexpected response: %s", action)
}

func (c *TradingClient) sendMessage(v interface{}) error {
	data, err := c.serializer.Marshal(v)
	if err != nil {
		return err
	}
	// 2 is websocket.BinaryMessage in gorilla, 1 is TextMessage
	msgType := 1
	if c.config.Encoding == "msgpack" {
		msgType = 2
	}
	return c.transport.WriteMessage(msgType, data)
}

func (c *TradingClient) receiveMessage(v interface{}) error {
	_, msg, err := c.transport.ReadMessage()
	if err != nil {
		return err
	}

	return c.serializer.Unmarshal(msg, v)
}

func (c *TradingClient) messageLoop(ctx context.Context) {
	defer close(c.done)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, msgBytes, err := c.transport.ReadMessage()
			if err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				return
			}

			var raw map[string]interface{}
			if mErr := c.serializer.Unmarshal(msgBytes, &raw); mErr != nil {
				if c.OnError != nil {
					c.OnError(fmt.Errorf("failed to unmarshal message: %w", mErr))
				}
				continue
			}

			// Pre-process Ping/Pong
			action, _ := raw["action"].(string)
			if action == "" {
				action, _ = raw["a"].(string)
			}
			
			if action == "ping" {
				_ = c.sendMessage(map[string]string{"action": "pong"})
				continue
			} else if action == "pong" || action == "subscribed" {
				continue
			}

			c.dispatcher.ProcessRawMessage(raw, msgBytes)
		}
	}
}

func (c *TradingClient) heartbeatLoop(ctx context.Context) {
	ticker := time.NewTicker(c.config.HeartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := c.sendMessage(map[string]string{"action": "ping"})
			if err != nil {
				log.Printf("failed to send heartbeat ping: %v", err)
			}
		}
	}
}

func (c *TradingClient) Disconnect() {
	if c.cancelFunc != nil {
		c.cancelFunc()
	}
	c.transport.Close()
}

func (c *TradingClient) Wait() {
	<-c.done
}
