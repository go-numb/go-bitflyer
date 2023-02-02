package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type BalanceHistory struct {
	CurrencyCode string `url:"currency_code,omitempty"`
	Count        int    `url:"count,omitempty"`
	Before       int64  `url:"before,omitempty"`
	After        int64  `url:"after,omitempty"`
}

type ResponseForBalanceHistory []BalanceHis

type BalanceHis struct {
	ID           int64   `json:"id"`
	TradeDate    string  `json:"trade_date"`
	EventDate    string  `json:"event_date"`
	ProductCode  string  `json:"product_code"`
	CurrencyCode string  `json:"currency_code"`
	TradeType    string  `json:"trade_type"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Quantity     float64 `json:"quantity"`
	Commission   float64 `json:"commission"`
	Balance      float64 `json:"balance"`
	OrderID      string  `json:"order_id"`
}

func (req *BalanceHistory) IsPrivate() bool {
	return true
}

func (req *BalanceHistory) Path() string {
	return "me/getbalancehistory"
}

func (req *BalanceHistory) Method() string {
	return http.MethodGet
}

func (req *BalanceHistory) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *BalanceHistory) Payload() []byte {
	return nil
}
