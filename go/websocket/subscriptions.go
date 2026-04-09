package websocket

import "fmt"

type ChannelConfig struct {
	Name    string   `json:"name" msgpack:"name"`
	Symbols []string `json:"symbols" msgpack:"symbols"`
}

type SubscriptionRequest struct {
	Action   string          `json:"action" msgpack:"action"`
	Channels []ChannelConfig `json:"channels" msgpack:"channels"`
}

var defaultBoards = []string{"G1", "G3", "G4", "G7", "T1", "T2", "T3", "T4", "T6"}

func (c *TradingClient) subscribeChannel(channelName string, symbols []string) error {
	if !c.isAuthenticated {
		return fmt.Errorf("must authenticate before subscribing")
	}

	req := SubscriptionRequest{
		Action: "subscribe",
		Channels: []ChannelConfig{
			{Name: channelName, Symbols: symbols},
		},
	}

	// Store for reconnection logic
	c.subsMutex.Lock()
	c.subscriptions[channelName] = req
	c.subsMutex.Unlock()

	return c.sendMessage(req)
}

func (c *TradingClient) SubscribeTrades(symbols []string, boardID string) error {
	boards := defaultBoards
	if boardID != "" {
		boards = []string{boardID}
	}
	for _, board := range boards {
		channel := fmt.Sprintf("tick.%s.%s", board, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeTradeExtra(symbols []string, boardID string) error {
	boards := defaultBoards
	if boardID != "" {
		boards = []string{boardID}
	}
	for _, board := range boards {
		channel := fmt.Sprintf("tick_extra.%s.%s", board, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeExpectedPrice(symbols []string, boardID string) error {
	boards := defaultBoards
	if boardID != "" {
		boards = []string{boardID}
	}
	for _, board := range boards {
		channel := fmt.Sprintf("expected_price.%s.%s", board, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeForeignInvestor(symbols []string, boardID string) error {
	boards := defaultBoards
	if boardID != "" {
		boards = []string{boardID}
	}
	for _, board := range boards {
		channel := fmt.Sprintf("foreign_investor.%s.%s", board, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeQuotes(symbols []string, boardID string) error {
	boards := []string{"G1", "G2", "G3", "G4", "G5", "G6", "G7"}
	if boardID != "" {
		boards = []string{boardID}
	}
	for _, board := range boards {
		channel := fmt.Sprintf("top_price.%s.%s", board, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeOHLC(symbols []string, resolution string) error {
	resolutions := []string{"1", "3", "5", "15", "30", "1H", "1D", "1W"}
	if resolution != "" {
		resolutions = []string{resolution}
	}

	for _, res := range resolutions {
		channel := fmt.Sprintf("ohlc.%s.%s", res, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeOHLCClosed(symbols []string, resolution string) error {
	resolutions := []string{"1", "3", "5", "15", "30", "1H", "1D", "1W"}
	if resolution != "" {
		resolutions = []string{resolution}
	}

	for _, res := range resolutions {
		channel := fmt.Sprintf("ohlc_closed.%s.%s", res, c.config.Encoding)
		if err := c.subscribeChannel(channel, symbols); err != nil {
			return err
		}
	}
	return nil
}

func (c *TradingClient) SubscribeMarketIndex(marketIndex string) error {
	channel := fmt.Sprintf("market_index.%s.%s", marketIndex, c.config.Encoding)
	return c.subscribeChannel(channel, []string{})
}

func (c *TradingClient) SubscribeOrders() error {
	return c.subscribeChannel("orders", []string{})
}

func (c *TradingClient) SubscribePositions() error {
	return c.subscribeChannel("positions", []string{})
}

func (c *TradingClient) SubscribeAccount() error {
	return c.subscribeChannel("account", []string{})
}

func (c *TradingClient) Unsubscribe(channelName string, symbols []string) error {
	req := SubscriptionRequest{
		Action: "unsubscribe",
		Channels: []ChannelConfig{
			{Name: channelName, Symbols: symbols},
		},
	}

	c.subsMutex.Lock()
	delete(c.subscriptions, channelName) // Simplistic map update logic
	c.subsMutex.Unlock()

	return c.sendMessage(req)
}
