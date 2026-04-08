package dnse

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"net/url"
	"strings"
)

// buildSignature generates the headers string and the signature for X-Signature
func buildSignature(secret string, method string, path string, dateValue string, algorithm string, nonce string, headerName string) (string, string) {
	headerKey := strings.ToLower(headerName)
	headers := fmt.Sprintf("(request-target) %s", headerKey)

	signatureString := fmt.Sprintf("(request-target): %s %s\n%s: %s", strings.ToLower(method), path, headerKey, dateValue)
	if nonce != "" {
		signatureString += fmt.Sprintf("\nnonce: %s", nonce)
	}

	var h func() hash.Hash
	switch algorithm {
	case "hmac-sha256":
		h = sha256.New
	case "hmac-sha384":
		h = sha512.New384
	case "hmac-sha512":
		h = sha512.New
	default: // using sha1 as fallback based on python code
		h = sha1.New
	}

	mac := hmac.New(h, []byte(secret))
	mac.Write([]byte(signatureString))
	encoded := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	escaped := url.QueryEscape(encoded)
	
	// Revert some characters if QueryEscape escapes too aggressively (like %2B for +)
	// Actually Python's `urllib.parse.quote(..., safe="")` escapes everything. 
	// Go's url.QueryEscape escapes space to + instead of %20. But base64 doesn't have spaces, it has + / =
	// Python's quote function escapes '+' to '%2B', '/' to '%2F', '=' to '%3D'. Go's url.QueryEscape does the same.

	return headers, escaped
}
