package websocket

import (
	"fmt"
)

// Dispatcher chịu trách nhiệm phân loại và kích hoạt các hàm callback
type Dispatcher struct {
	serializer Serializer
	client     *TradingClient // Tham chiếu ngược về Client giữ các callback struct để giữ API tương thích ngược
}

// NewDispatcher tạo router sự kiện
func NewDispatcher(client *TradingClient, s Serializer) *Dispatcher {
	return &Dispatcher{
		client:     client,
		serializer: s,
	}
}

// ProcessRawMessage phân tách json/msgpack object ra action/type và gọi callback tương ứng
func (d *Dispatcher) ProcessRawMessage(raw map[string]interface{}, rawBytes []byte) error {
	action, _ := raw["action"].(string)
	if action == "" {
		action, _ = raw["a"].(string)
	}

	if action == "error" {
		msg, _ := raw["message"].(string)
		if d.client.OnError != nil {
			d.client.OnError(fmt.Errorf("server error: %s", msg))
		}
		return nil
	} else if action != "" {
		// Các system action "auth_success", "ping", "pong" xử lý ở Transport/Client wrapper.
		// Ở đây chỉ nhận data events
		return nil
	}

	msgType, _ := raw["T"].(string)

	switch msgType {
	case "t":
		if d.client.OnTrade != nil {
			var obj Trade
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnTrade(obj)
			}
		}
	case "te":
		if d.client.OnTradeExtra != nil {
			var obj TradeExtra
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnTradeExtra(obj)
			}
		}
	case "e":
		if d.client.OnExpectedPrice != nil {
			var obj ExpectedPrice
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnExpectedPrice(obj)
			}
		}
	case "sd":
		if d.client.OnSecurityDefinition != nil {
			var obj SecurityDefinition
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnSecurityDefinition(obj)
			}
		}
	case "q":
		if d.client.OnQuote != nil {
			var obj Quote
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnQuote(obj)
			}
		}
	case "b":
		if d.client.OnOhlc != nil {
			var obj Ohlc
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnOhlc(obj)
			}
		}
	case "bc":
		if d.client.OnOhlcClosed != nil {
			var obj Ohlc
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnOhlcClosed(obj)
			}
		}
	case "o":
		if d.client.OnOrder != nil {
			var obj Order
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnOrder(obj)
			}
		}
	case "p":
		if d.client.OnPosition != nil {
			var obj Position
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnPosition(obj)
			}
		}
	case "a":
		if d.client.OnAccountUpdate != nil {
			var obj AccountUpdate
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnAccountUpdate(obj)
			}
		}
	case "mi":
		if d.client.OnMarketIndex != nil {
			var obj MarketIndex
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnMarketIndex(obj)
			}
		}
	case "f":
		if d.client.OnForeignInvestor != nil {
			var obj ForeignInvestor
			if err := d.serializer.Unmarshal(rawBytes, &obj); err == nil {
				d.client.OnForeignInvestor(obj)
			}
		}
	}

	return nil
}
