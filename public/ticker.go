package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Ticker struct {
	ProductCode string `url:"product_code,omitempty"`
}

type ResponseForTicker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int64   `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
	PreopenEnd      string  `json:"preopen_end"`
	CircuitBreakEnd string  `json:"circuit_break_end"`
}

func (req *Ticker) IsPrivate() bool {
	return false
}

func (req *Ticker) Path() string {
	return "ticker"
}

func (req *Ticker) Method() string {
	return http.MethodGet
}

func (req *Ticker) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Ticker) Payload() []byte {
	return nil
}
