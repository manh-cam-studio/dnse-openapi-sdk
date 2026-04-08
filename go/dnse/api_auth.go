package dnse

// CreateTradingToken generates a Trading Token required for order placement.
func (c *Client) CreateTradingToken(otpType, passcode string, dryRun bool) (int, []byte, error) {
	body := map[string]string{
		"otpType":  otpType,
		"passcode": passcode,
	}
	return c.Request("POST", "/registration/trading-token", nil, body, nil, dryRun)
}

// SendEmailOTP requests an OTP sent to your registered email.
func (c *Client) SendEmailOTP(dryRun bool) (int, []byte, error) {
	return c.Request("POST", "/registration/send-email-otp", nil, nil, nil, dryRun)
}
