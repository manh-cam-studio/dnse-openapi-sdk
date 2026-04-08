package dnse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	APIKey           string
	APISecret        string
	BaseURL          string
	Algorithm        string
	HMACNonceEnabled bool
	HTTPClient       *http.Client
	Debug            bool
}

type Client struct {
	config Config
}

func NewClient(apiKey, apiSecret string) *Client {
	return NewClientWithConfig(Config{
		APIKey:           apiKey,
		APISecret:        apiSecret,
		BaseURL:          "https://openapi.dnse.com.vn",
		Algorithm:        "hmac-sha256",
		HMACNonceEnabled: true,
	})
}

func NewClientWithConfig(cfg Config) *Client {
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://openapi.dnse.com.vn"
	}
	if cfg.Algorithm == "" {
		cfg.Algorithm = "hmac-sha256"
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: 60 * time.Second,
		}
	}
	return &Client{config: cfg}
}

func (c *Client) getDateHeader() string {
	val := os.Getenv("DATE_HEADER")
	if val == "" {
		val = "Date"
	}
	return val
}

func (c *Client) signatureHeaders(method, path string) (string, string) {
	dateValue := time.Now().UTC().Format(time.RFC1123)
	
	var nonce string
	if c.config.HMACNonceEnabled {
		nonceHex := uuid.New().String()
		nonce = string(bytes.ReplaceAll([]byte(nonceHex), []byte("-"), []byte(""))) // hex only like in uuid.uuid4().hex
	}

	headersList, signature := buildSignature(
		c.config.APISecret,
		method,
		path,
		dateValue,
		c.config.Algorithm,
		string(nonce),
		c.getDateHeader(),
	)

	sigHeaderVal := fmt.Sprintf(`Signature keyId="%s",algorithm="%s",headers="%s",signature="%s"`,
		c.config.APIKey, c.config.Algorithm, headersList, signature)

	if string(nonce) != "" {
		sigHeaderVal += fmt.Sprintf(`,nonce="%s"`, string(nonce))
	}

	return dateValue, sigHeaderVal
}

// Request performs a REST API request and handles HMAC signatures
func (c *Client) Request(method, path string, query map[string]string, body interface{}, headers map[string]string, dryRun bool) (int, []byte, error) {
	debug := c.config.Debug || os.Getenv("DEBUG") == "true"

	urlStr := c.config.BaseURL + path

	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return 0, nil, err
	}

	if len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	var data []byte
	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return 0, nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(data))
		req.ContentLength = int64(len(data))
	}

	dateValue, sigHeaderValue := c.signatureHeaders(method, req.URL.Path)
	dateHeaderName := c.getDateHeader()

	req.Header.Set(dateHeaderName, dateValue)
	req.Header.Set("X-Signature", sigHeaderValue)
	req.Header.Set("x-api-key", c.config.APIKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if dryRun || debug {
		prefix := "DEBUG"
		if dryRun {
			prefix = "DRY RUN"
		}
		log.Printf("%s url: %s", prefix, req.URL.String())
		log.Printf("%s method: %s", prefix, method)
		log.Printf("%s query: %v", prefix, query)
		log.Printf("%s headers: %v", prefix, req.Header)
		if body != nil {
			log.Printf("%s body: %s", prefix, string(data))
		}
	}

	if dryRun {
		return 200, nil, nil
	}

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, respBody, nil
}
