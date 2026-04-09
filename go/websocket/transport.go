package websocket

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Transport định nghĩa giao diện giao tiếp thời gian thực
type Transport interface {
	Connect(ctx context.Context, url string, timeout time.Duration) error
	ReadMessage() (int, []byte, error)
	WriteMessage(messageType int, data []byte) error
	Close() error
	IsConnected() bool
}

// GorillaTransport thực thi Transport interface dựa trên thư viện Gorilla Websocket
type GorillaTransport struct {
	conn      *websocket.Conn
	connMutex sync.Mutex
}

func NewGorillaTransport() *GorillaTransport {
	return &GorillaTransport{}
}

func (t *GorillaTransport) Connect(ctx context.Context, url string, timeout time.Duration) error {
	dialer := websocket.Dialer{
		HandshakeTimeout: timeout,
	}

	conn, _, err := dialer.DialContext(ctx, url, nil)
	if err != nil {
		return fmt.Errorf("transport failed to connect: %w", err)
	}

	t.connMutex.Lock()
	t.conn = conn
	t.connMutex.Unlock()

	return nil
}

func (t *GorillaTransport) ReadMessage() (int, []byte, error) {
	t.connMutex.Lock()
	conn := t.conn
	t.connMutex.Unlock()

	if conn == nil {
		return 0, nil, errors.New("transport not connected")
	}

	return conn.ReadMessage()
}

func (t *GorillaTransport) WriteMessage(messageType int, data []byte) error {
	t.connMutex.Lock()
	defer t.connMutex.Unlock()

	if t.conn == nil {
		return errors.New("transport not connected")
	}

	return t.conn.WriteMessage(messageType, data)
}

func (t *GorillaTransport) Close() error {
	t.connMutex.Lock()
	defer t.connMutex.Unlock()

	if t.conn != nil {
		err := t.conn.Close()
		t.conn = nil
		return err
	}
	return nil
}

func (t *GorillaTransport) IsConnected() bool {
	t.connMutex.Lock()
	defer t.connMutex.Unlock()
	return t.conn != nil
}
