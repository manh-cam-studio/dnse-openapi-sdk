package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
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

	conn       *websocket.Conn
	connMutex  sync.Mutex
	cancelFunc context.CancelFunc

	isAuthenticated bool
	sessionID       string

	done chan struct{}

	subscriptions map[string]SubscriptionRequest

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
	return &TradingClient{
		config:        cfg,
		subscriptions: make(map[string]SubscriptionRequest),
		done:          make(chan struct{}),
	}
}

func (c *TradingClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	c.cancelFunc = cancel

	url := fmt.Sprintf("%s/v1/stream?encoding=%s", c.config.BaseURL, c.config.Encoding)

	dialer := websocket.Dialer{
		HandshakeTimeout: c.config.Timeout,
	}

	conn, _, err := dialer.DialContext(ctx, url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.connMutex.Lock()
	c.conn = conn
	c.connMutex.Unlock()

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
	c.connMutex.Lock()
	defer c.connMutex.Unlock()

	if c.conn == nil {
		return errors.New("not connected")
	}

	if c.config.Encoding == "msgpack" {
		data, err := msgpack.Marshal(v)
		if err != nil {
			return err
		}
		return c.conn.WriteMessage(websocket.BinaryMessage, data)
	} else {
		return c.conn.WriteJSON(v)
	}
}

func (c *TradingClient) receiveMessage(v interface{}) error {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}

	if c.config.Encoding == "msgpack" {
		return msgpack.Unmarshal(msg, v)
	} else {
		return json.Unmarshal(msg, v)
	}
}

func (c *TradingClient) messageLoop(ctx context.Context) {
	defer close(c.done)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				// Handle reconnection logic in production...
				return
			}

			var raw map[string]interface{}
			var mErr error
			if c.config.Encoding == "msgpack" {
				mErr = msgpack.Unmarshal(msg, &raw)
			} else {
				mErr = json.Unmarshal(msg, &raw)
			}
			
			if mErr != nil {
				if c.OnError != nil {
					c.OnError(fmt.Errorf("failed to unmarshal message: %w", mErr))
				}
				continue
			}

			c.dispatchMessage(raw, msg)
		}
	}
}

func (c *TradingClient) dispatchMessage(raw map[string]interface{}, rawMsg []byte) {
	action, _ := raw["action"].(string)
	if action == "" {
		action, _ = raw["a"].(string)
	}

	if action == "ping" {
		_ = c.sendMessage(map[string]string{"action": "pong"})
		return
	} else if action == "pong" || action == "subscribed" {
		return
	} else if action == "error" {
		msg, _ := raw["message"].(string)
		if c.OnError != nil {
			c.OnError(fmt.Errorf("server error: %s", msg))
		}
		return
	}

	msgType, _ := raw["T"].(string)

	var unmarshal func(interface{}) error
	if c.config.Encoding == "msgpack" {
		unmarshal = func(v interface{}) error { return msgpack.Unmarshal(rawMsg, v) }
	} else {
		unmarshal = func(v interface{}) error { return json.Unmarshal(rawMsg, v) }
	}

	switch msgType {
	case "t":
		if c.OnTrade != nil {
			var obj Trade
			_ = unmarshal(&obj)
			c.OnTrade(obj)
		}
	case "te":
		if c.OnTradeExtra != nil {
			var obj TradeExtra
			_ = unmarshal(&obj)
			c.OnTradeExtra(obj)
		}
	case "e":
		if c.OnExpectedPrice != nil {
			var obj ExpectedPrice
			_ = unmarshal(&obj)
			c.OnExpectedPrice(obj)
		}
	case "sd":
		if c.OnSecurityDefinition != nil {
			var obj SecurityDefinition
			_ = unmarshal(&obj)
			c.OnSecurityDefinition(obj)
		}
	case "q":
		if c.OnQuote != nil {
			var obj Quote
			_ = unmarshal(&obj)
			c.OnQuote(obj)
		}
	case "b":
		if c.OnOhlc != nil {
			var obj Ohlc
			_ = unmarshal(&obj)
			c.OnOhlc(obj)
		}
	case "bc":
		if c.OnOhlcClosed != nil {
			var obj Ohlc
			_ = unmarshal(&obj)
			c.OnOhlcClosed(obj)
		}
	case "o":
		if c.OnOrder != nil {
			var obj Order
			_ = unmarshal(&obj)
			c.OnOrder(obj)
		}
	case "p":
		if c.OnPosition != nil {
			var obj Position
			_ = unmarshal(&obj)
			c.OnPosition(obj)
		}
	case "a":
		if c.OnAccountUpdate != nil {
			var obj AccountUpdate
			_ = unmarshal(&obj)
			c.OnAccountUpdate(obj)
		}
	case "mi":
		if c.OnMarketIndex != nil {
			var obj MarketIndex
			_ = unmarshal(&obj)
			c.OnMarketIndex(obj)
		}
	case "f":
		if c.OnForeignInvestor != nil {
			var obj ForeignInvestor
			_ = unmarshal(&obj)
			c.OnForeignInvestor(obj)
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
	c.connMutex.Lock()
	defer c.connMutex.Unlock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}

func (c *TradingClient) Wait() {
	<-c.done
}
