package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Markets struct {
}

type ResponseForMarkets []Market

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
}

func (req *Markets) IsPrivate() bool {
	return false
}

func (req *Markets) Path() string {
	return "markets"
}

func (req *Markets) Method() string {
	return http.MethodGet
}

func (req *Markets) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Markets) Payload() []byte {
	return nil
}
