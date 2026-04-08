package websocket

import (
	"strconv"
)

type PriceLevel struct {
	Price    float64   `json:"price" msgpack:"price"`
	Quantity FlexInt64 `json:"qtty" msgpack:"qtty"`
}

type Trade struct {
	MarketID          string    `json:"marketId" msgpack:"marketId"`
	BoardID           string    `json:"boardId" msgpack:"boardId"`
	Isin              string    `json:"isin" msgpack:"isin"`
	Symbol            string    `json:"symbol" msgpack:"symbol"`
	Price             float64   `json:"matchPrice" msgpack:"matchPrice"`
	Quantity          FlexInt64 `json:"matchQtty" msgpack:"matchQtty"`
	TotalVolumeTraded FlexInt64 `json:"totalVolumeTraded" msgpack:"totalVolumeTraded"`
	GrossTradeAmount  float64   `json:"grossTradeAmount" msgpack:"grossTradeAmount"`
	HighestPrice      float64   `json:"highestPrice" msgpack:"highestPrice"`
	LowestPrice       float64   `json:"lowestPrice" msgpack:"lowestPrice"`
	OpenPrice         float64   `json:"openPrice" msgpack:"openPrice"`
	TradingSessionID  string    `json:"tradingSessionId" msgpack:"tradingSessionId"`
	ReceivedAt        float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type TradeExtra struct {
	MarketID          string    `json:"marketId" msgpack:"marketId"`
	BoardID           string    `json:"boardId" msgpack:"boardId"`
	Isin              string    `json:"isin" msgpack:"isin"`
	Symbol            string    `json:"symbol" msgpack:"symbol"`
	Price             float64   `json:"matchPrice" msgpack:"matchPrice"`
	Quantity          FlexInt64 `json:"matchQtty" msgpack:"matchQtty"`
	Side              string    `json:"side" msgpack:"side"` // UNSPECIFIED, B, S
	AvgPrice          float64   `json:"avgPrice" msgpack:"avgPrice"`
	TotalVolumeTraded FlexInt64 `json:"totalVolumeTraded" msgpack:"totalVolumeTraded"`
	GrossTradeAmount  float64   `json:"grossTradeAmount" msgpack:"grossTradeAmount"`
	HighestPrice      float64   `json:"highestPrice" msgpack:"highestPrice"`
	LowestPrice       float64   `json:"lowestPrice" msgpack:"lowestPrice"`
	OpenPrice         float64   `json:"openPrice" msgpack:"openPrice"`
	TradingSessionID  string    `json:"tradingSessionId" msgpack:"tradingSessionId"`
	ReceivedAt        float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type ForeignInvestor struct {
	MarketID                     string    `json:"marketId" msgpack:"marketId"`
	BoardID                      string    `json:"boardId" msgpack:"boardId"`
	TradingSessionID             string    `json:"tradingSessionId" msgpack:"tradingSessionId"`
	Symbol                       string    `json:"symbol" msgpack:"symbol"`
	TransactTime                 string    `json:"transactTime" msgpack:"transactTime"`
	ForeignInvestorTypeCode      string    `json:"foreignInvestorTypeCode" msgpack:"foreignInvestorTypeCode"`
	SellVolume                   FlexInt64 `json:"sellVolume" msgpack:"sellVolume"`
	SellTradedAmount             FlexInt64 `json:"sellTradedAmount" msgpack:"sellTradedAmount"`
	BuyVolume                    FlexInt64 `json:"buyVolume" msgpack:"buyVolume"`
	BuyTradedAmount              FlexInt64 `json:"buyTradedAmount" msgpack:"buyTradedAmount"`
	TotalSellVolume              FlexInt64 `json:"totalSellVolume" msgpack:"totalSellVolume"`
	TotalSellTradedAmount        FlexInt64 `json:"totalSellTradedAmount" msgpack:"totalSellTradedAmount"`
	TotalBuyVolume               FlexInt64 `json:"totalBuyVolume" msgpack:"totalBuyVolume"`
	TotalBuyTradedAmount         FlexInt64 `json:"totalBuyTradedAmount" msgpack:"totalBuyTradedAmount"`
	ForeignerOrderLimitQuantity  FlexInt64 `json:"foreignerOrderLimitQuantity" msgpack:"foreignerOrderLimitQuantity"`
	ForeignerBuyPossibleQuantity FlexInt64 `json:"foreignerBuyPossibleQuantity" msgpack:"foreignerBuyPossibleQuantity"`
	ReceivedAt                   float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type MarketIndex struct {
	IndexName                        string    `json:"indexName" msgpack:"indexName"`
	ChangedRatio                     float64   `json:"changedRatio" msgpack:"changedRatio"`
	ChangedValue                     float64   `json:"changedValue" msgpack:"changedValue"`
	FluctuationSteadinessIssueCount  int       `json:"fluctuationSteadinessIssueCount" msgpack:"fluctuationSteadinessIssueCount"`
	FluctuationDownIssueCount        int       `json:"fluctuationDownIssueCount" msgpack:"fluctuationDownIssueCount"`
	FluctuationUpIssueCount          int       `json:"fluctuationUpIssueCount" msgpack:"fluctuationUpIssueCount"`
	FluctuationLowerLimitIssueCount  int       `json:"fluctuationLowerLimitIssueCount" msgpack:"fluctuationLowerLimitIssueCount"`
	FluctuationUpperLimitIssueCount  int       `json:"fluctuationUpperLimitIssueCount" msgpack:"fluctuationUpperLimitIssueCount"`
	FluctuationDownIssueVolume       FlexInt64 `json:"fluctuationDownIssueVolume" msgpack:"fluctuationDownIssueVolume"`
	FluctuationUpIssueVolume         FlexInt64 `json:"fluctuationUpIssueVolume" msgpack:"fluctuationUpIssueVolume"`
	FluctuationSteadinessIssueVolume FlexInt64 `json:"fluctuationSteadinessIssueVolume" msgpack:"fluctuationSteadinessIssueVolume"`
	CurrencyCode                     string    `json:"currencyCode" msgpack:"currencyCode"`
	IndexTypeCode                    string    `json:"indexTypeCode" msgpack:"indexTypeCode"`
	LowestValueIndexes               float64   `json:"lowestValueIndexes" msgpack:"lowestValueIndexes"`
	HighestValueIndexes              float64   `json:"highestValueIndexes" msgpack:"highestValueIndexes"`
	PriorValueIndexes                float64   `json:"priorValueIndexes" msgpack:"priorValueIndexes"`
	ValueIndexes                     float64   `json:"valueIndexes" msgpack:"valueIndexes"`
	ContauctAccTrdVal                float64   `json:"contauctAccTrdVal" msgpack:"contauctAccTrdVal"`
	ContauctAccTrdVol                FlexInt64 `json:"contauctAccTrdVol" msgpack:"contauctAccTrdVol"`
	BlkTrdAccTrdVal                  float64   `json:"blkTrdAccTrdVal" msgpack:"blkTrdAccTrdVal"`
	BlkTrdAccTrdVol                  FlexInt64 `json:"blkTrdAccTrdVol" msgpack:"blkTrdAccTrdVol"`
	GrossTradeAmount                 float64   `json:"grossTradeAmount" msgpack:"grossTradeAmount"`
	TotalVolumeTraded                FlexInt64 `json:"totalVolumeTraded" msgpack:"totalVolumeTraded"`
	MarketIndexClass                 int       `json:"marketIndexClass" msgpack:"marketIndexClass"`
	MarketID                         string    `json:"marketId" msgpack:"marketId"`
	TradingSessionID                 string    `json:"tradingSessionId" msgpack:"tradingSessionId"`
	TransactTime                     string    `json:"transactTime" msgpack:"transactTime"` 
	ReceivedAt                       float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type ExpectedPrice struct {
	MarketID              string    `json:"marketId" msgpack:"marketId"`
	BoardID               string    `json:"boardId" msgpack:"boardId"`
	Isin                  string    `json:"isin" msgpack:"isin"`
	Symbol                string    `json:"symbol" msgpack:"symbol"`
	ClosePrice            float64   `json:"closePrice" msgpack:"closePrice"`
	ExpectedTradePrice    float64   `json:"expectedTradePrice" msgpack:"expectedTradePrice"`
	ExpectedTradeQuantity FlexInt64 `json:"expectedTradeQuantity" msgpack:"expectedTradeQuantity"`
	ReceivedAt            float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type SecurityDefinition struct {
	MarketID                        string    `json:"marketId" msgpack:"marketId"`
	BoardID                         string    `json:"boardId" msgpack:"boardId"`
	Symbol                          string    `json:"symbol" msgpack:"symbol"`
	Isin                            string    `json:"isin" msgpack:"isin"`
	ProductGrpID                    string    `json:"productGrpId" msgpack:"productGrpId"`
	SecurityGroupID                 string    `json:"securityGroupId" msgpack:"securityGroupId"`
	BasicPrice                      float64   `json:"basicPrice" msgpack:"basicPrice"`
	CeilingPrice                    float64   `json:"ceilingPrice" msgpack:"ceilingPrice"`
	FloorPrice                      float64   `json:"floorPrice" msgpack:"floorPrice"`
	OpenInterestQuantity            FlexInt64 `json:"openInterestQuantity" msgpack:"openInterestQuantity"`
	SecurityStatus                  string    `json:"securityStatus" msgpack:"securityStatus"`
	SymbolAdminStatusCode           string    `json:"symbolAdminStatusCode" msgpack:"symbolAdminStatusCode"`
	SymbolTradingMethodStatusCode   string    `json:"symbolTradingMethodStatusCode" msgpack:"symbolTradingMethodStatusCode"`
	SymbolTradingSanctionStatusCode string    `json:"symbolTradingSanctionStatusCode" msgpack:"symbolTradingSanctionStatusCode"`
	ReceivedAt                      float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type Quote struct {
	MarketID       string       `json:"marketId" msgpack:"marketId"`
	BoardID        string       `json:"boardId" msgpack:"boardId"`
	Symbol         string       `json:"symbol" msgpack:"symbol"`
	Isin           string       `json:"isin" msgpack:"isin"`
	Bid            []PriceLevel `json:"bid" msgpack:"bid"`
	Offer          []PriceLevel `json:"offer" msgpack:"offer"`
	TotalOfferQtty float64      `json:"totalOfferQtty" msgpack:"totalOfferQtty"`
	TotalBidQtty   float64      `json:"totalBidQtty" msgpack:"totalBidQtty"`
	ReceivedAt     float64      `json:"_receivedAt" msgpack:"_receivedAt"`
}

type Ohlc struct {
	Symbol      string    `json:"symbol" msgpack:"symbol"`
	Resolution  string    `json:"resolution" msgpack:"resolution"`
	Open        float64   `json:"open" msgpack:"open"`
	High        float64   `json:"high" msgpack:"high"`
	Low         float64   `json:"low" msgpack:"low"`
	Close       float64   `json:"close" msgpack:"close"`
	Volume      FlexInt64 `json:"volume" msgpack:"volume"`
	Time        FlexInt64 `json:"time" msgpack:"time"`
	Type        string    `json:"type" msgpack:"type"`
	LastUpdated FlexInt64 `json:"lastUpdated" msgpack:"lastUpdated"`
	ReceivedAt  float64   `json:"_receivedAt" msgpack:"_receivedAt"`
}

type Order struct {
	ID               string    `json:"id" msgpack:"id"`
	Side             string    `json:"side" msgpack:"side"`
	AccountNo        string    `json:"accountNo" msgpack:"accountNo"`
	Symbol           string    `json:"symbol" msgpack:"symbol"`
	Price            float64   `json:"price" msgpack:"price"`
	PriceSecure      float64   `json:"priceSecure" msgpack:"priceSecure"`
	AveragePrice     float64   `json:"averagePrice" msgpack:"averagePrice"`
	Quantity         FlexInt64 `json:"quantity" msgpack:"quantity"`
	FillQuantity     FlexInt64 `json:"fillQuantity" msgpack:"fillQuantity"`
	CanceledQuantity FlexInt64 `json:"canceledQuantity" msgpack:"canceledQuantity"`
	LeaveQuantity    FlexInt64 `json:"leaveQuantity" msgpack:"leaveQuantity"`
	OrderType        string    `json:"orderType" msgpack:"orderType"`
	OrderStatus      string    `json:"orderStatus" msgpack:"orderStatus"`
	LoanPackageID    FlexInt64 `json:"loanPackageId" msgpack:"loanPackageId"`
	MarketType       string    `json:"marketType" msgpack:"marketType"`
	TransDate        string    `json:"transDate" msgpack:"transDate"`
	CreatedDate      string    `json:"createdDate" msgpack:"createdDate"`
	ModifiedDate     string    `json:"modifiedDate" msgpack:"modifiedDate"`
}

type Position struct {
	Symbol              string    `json:"symbol" msgpack:"symbol"`
	Quantity            FlexInt64 `json:"quantity" msgpack:"quantity"`
	AveragePrice        string    `json:"averagePrice" msgpack:"averagePrice"`
	MarketValue         string    `json:"marketValue" msgpack:"marketValue"`
	CostBasis           string    `json:"costBasis" msgpack:"costBasis"`
	UnrealizedPl        string    `json:"unrealizedPl" msgpack:"unrealizedPl"`
	UnrealizedPlPercent string    `json:"unrealizedPlPercent" msgpack:"unrealizedPlPercent"`
	Timestamp           string    `json:"timestamp" msgpack:"timestamp"` 
}

type AccountUpdate struct {
	Cash           string `json:"cash" msgpack:"cash"`
	BuyingPower    string `json:"buyingPower" msgpack:"buyingPower"`
	PortfolioValue string `json:"portfolioValue" msgpack:"portfolioValue"`
	Equity         string `json:"equity" msgpack:"equity"`
	Timestamp      string `json:"timestamp" msgpack:"timestamp"` 
}

// Ensure interface compatibility format print
func (fi FlexInt64) String() string {
	return strconv.FormatInt(int64(fi), 10)
}
