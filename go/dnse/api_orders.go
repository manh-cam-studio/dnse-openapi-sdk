package dnse

import (
	"fmt"
	"strconv"
)

// GetOrders retrieves intraday order book.
func (c *Client) GetOrders(accountNo, marketType, orderCategory string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory != "" {
		query["orderCategory"] = orderCategory
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/orders", accountNo), query, nil, nil, dryRun)
}

// GetOrderDetail retrieves detailed information of a specific order (by ID).
func (c *Client) GetOrderDetail(accountNo, orderID, marketType, orderCategory string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory != "" {
		query["orderCategory"] = orderCategory
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/orders/%s", accountNo, orderID), query, nil, nil, dryRun)
}

// GetExecutionDetail retrieves derivative execution details and reports for a specific order.
func (c *Client) GetExecutionDetail(accountNo, orderID, marketType, orderCategory string, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory == "" {
		orderCategory = "NORMAL"
	}
	query["orderCategory"] = orderCategory
	return c.Request("GET", fmt.Sprintf("/accounts/%s/executions/%s", accountNo, orderID), query, nil, nil, dryRun)
}

// GetOrderHistory retrieves historical orders.
func (c *Client) GetOrderHistory(accountNo, marketType, fromDate, toDate string, pageSize, pageIndex *int, dryRun bool) (int, []byte, error) {
	query := map[string]string{
		"marketType": marketType,
	}
	if fromDate != "" {
		query["from"] = fromDate
	}
	if toDate != "" {
		query["to"] = toDate
	}
	if pageSize != nil {
		query["pageSize"] = strconv.Itoa(*pageSize)
	}
	if pageIndex != nil {
		query["pageIndex"] = strconv.Itoa(*pageIndex)
	}
	return c.Request("GET", fmt.Sprintf("/accounts/%s/orders/history", accountNo), query, nil, nil, dryRun)
}

// PostOrder submits a new trading order.
func (c *Client) PostOrder(marketType string, payload interface{}, tradingToken, orderCategory string, dryRun bool) (int, []byte, error) {
	headers := map[string]string{
		"trading-token": tradingToken,
	}
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory == "" {
		orderCategory = "NORMAL"
	}
	query["orderCategory"] = orderCategory
	return c.Request("POST", "/accounts/orders", query, payload, headers, dryRun)
}

// PutOrder modifies an existing order.
func (c *Client) PutOrder(accountNo, orderID, marketType string, payload interface{}, tradingToken, orderCategory string, dryRun bool) (int, []byte, error) {
	headers := map[string]string{
		"trading-token": tradingToken,
	}
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory != "" {
		query["orderCategory"] = orderCategory
	}
	return c.Request("PUT", fmt.Sprintf("/accounts/%s/orders/%s", accountNo, orderID), query, payload, headers, dryRun)
}

// CancelOrder cancels an existing order.
func (c *Client) CancelOrder(accountNo, orderID, marketType, tradingToken, orderCategory string, dryRun bool) (int, []byte, error) {
	headers := map[string]string{
		"trading-token": tradingToken,
	}
	query := map[string]string{
		"marketType": marketType,
	}
	if orderCategory != "" {
		query["orderCategory"] = orderCategory
	}
	return c.Request("DELETE", fmt.Sprintf("/accounts/%s/orders/%s", accountNo, orderID), query, nil, headers, dryRun)
}
