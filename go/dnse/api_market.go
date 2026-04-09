package dnse

import (
	"fmt"
	"strconv"
)

// GetSecurityDefinition retrieves securities definition updates.
func (c *Client) GetSecurityDefinition(symbol, boardID string, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	if boardID != "" {
		query["boardId"] = boardID
	}
	return c.Request("GET", fmt.Sprintf("/price/%s/secdef", symbol), query, nil, nil, dryRun)
}

// GetOHLC retrieves OHLC data for a specific instrument over a given time range.
func (c *Client) GetOHLC(barType string, queryParams map[string]string, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	for k, v := range queryParams {
		query[k] = v
	}
	query["type"] = barType
	return c.Request("GET", "/price/ohlc", query, nil, nil, dryRun)
}

// GetTrades retrieves historical trade data for a specific instrument.
func (c *Client) GetTrades(symbol, boardID, fromDate, toDate string, limit *int, order, nextPageToken string, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	if boardID != "" {
		query["boardId"] = boardID
	}
	if fromDate != "" {
		query["from"] = fromDate
	}
	if toDate != "" {
		query["to"] = toDate
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}
	if order != "" {
		query["order"] = order
	}
	if nextPageToken != "" {
		query["nextPageToken"] = nextPageToken
	}
	return c.Request("GET", fmt.Sprintf("/price/%s/trades", symbol), query, nil, nil, dryRun)
}

// GetInstruments retrieves the list of available trading instruments and their metadata.
func (c *Client) GetInstruments(symbol, marketID, securityGroupID, indexName string, limit, page *int, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	if symbol != "" {
		query["symbol"] = symbol
	}
	if marketID != "" {
		query["marketId"] = marketID
	}
	if securityGroupID != "" {
		query["securityGroupId"] = securityGroupID
	}
	if indexName != "" {
		query["indexName"] = indexName
	}
	if limit != nil {
		query["limit"] = strconv.Itoa(*limit)
	}
	if page != nil {
		query["page"] = strconv.Itoa(*page)
	}
	return c.Request("GET", "/instruments", query, nil, nil, dryRun)
}

// GetLatestTrade retrieves the most recent trade for a specific instrument.
func (c *Client) GetLatestTrade(symbol, boardID string, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	if boardID != "" {
		query["boardId"] = boardID
	}
	return c.Request("GET", fmt.Sprintf("/price/%s/trades/latest", symbol), query, nil, nil, dryRun)
}

// GetClosePrice retrieves the close price.
func (c *Client) GetClosePrice(symbol, boardID string, dryRun bool) (int, []byte, error) {
	query := make(map[string]string)
	if boardID != "" {
		query["boardId"] = boardID
	}
	return c.Request("GET", fmt.Sprintf("/price/%s/close", symbol), query, nil, nil, dryRun)
}
