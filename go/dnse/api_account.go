package dnse

import (
	"fmt"
)

// GetAccounts retrieves all trading sub-accounts managed under the account corresponding to the API Key.
func (c *Client) GetAccounts(dryRun bool) (int, []byte, error) {
	return c.Request("GET", "/accounts", nil, nil, nil, dryRun)
}

// GetBalances retrieves asset balances of a trading sub-account.
func (c *Client) GetBalances(accountNo string, dryRun bool) (int, []byte, error) {
	return c.Request("GET", fmt.Sprintf("/accounts/%s/balances", accountNo), nil, nil, nil, dryRun)
}

// GetLoanPackages retrieves available loan package codes.
func (c *Client) GetLoanPackages(accountNo, marketType, symbol string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	if symbol != "" {
		query["symbol"] = symbol
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/loan-packages", accountNo), query, nil, nil, dryRun)
}

// GetPositions retrieves current holding positions.
func (c *Client) GetPositions(accountNo, marketType string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/positions", accountNo), query, nil, nil, dryRun)
}

// GetPositionByID retrieves detailed information of a specific position.
func (c *Client) GetPositionByID(marketType, positionID string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	return c.Request("GET", fmt.Sprintf("/accounts/positions/%s", positionID), query, nil, nil, dryRun)
}

// GetPPSE retrieves buying power and selling power before placing an order.
func (c *Client) GetPPSE(accountNo, marketType, symbol string, price float64, loanPackageID string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType":    marketType,
		"symbol":        symbol,
		"price":         fmt.Sprintf("%v", price),
		"loanPackageId": loanPackageID,
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/ppse", accountNo), query, nil, nil, dryRun)
}

// ClosePosition closes an existing position (by ID).
func (c *Client) ClosePosition(positionID, marketType, tradingToken string, dryRun bool) (int, []byte, error) {
	headers := map[string]string{
		"trading-token": tradingToken,
	}
	query := map[string]string{
		"marketType": marketType,
	}
	return c.Request("POST", fmt.Sprintf("/accounts/positions/%s/close", positionID), query, nil, headers, dryRun)
}
