package websocket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type AuthMessage struct {
	Action    string `json:"action" msgpack:"action"`
	APIKey    string `json:"api_key" msgpack:"api_key"`
	Signature string `json:"signature" msgpack:"signature"`
	Timestamp int64  `json:"timestamp" msgpack:"timestamp"`
	Nonce     string `json:"nonce" msgpack:"nonce"`
}

// CreateAuthMessage generates the authentication payload needed for the websocket connection
func CreateAuthMessage(apiKey, apiSecret string) AuthMessage {
	now := time.Now()
	timestamp := now.Unix() // second
	nonce := fmt.Sprintf("%d", now.UnixMicro())

	signature := ComputeSignature(apiSecret, apiKey, timestamp, nonce)

	return AuthMessage{
		Action:    "auth",
		APIKey:    apiKey,
		Signature: signature,
		Timestamp: timestamp,
		Nonce:     nonce,
	}
}

// ComputeSignature calculates HMAC-SHA256 signature
func ComputeSignature(apiSecret, apiKey string, timestamp int64, nonce string) string {
	// Message format: {api_key}:{timestamp}:{nonce}
	message := fmt.Sprintf("%s:%d:%s", apiKey, timestamp, nonce)

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
