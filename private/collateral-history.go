package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type CollateralHistory struct {
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

type ResponseForCollateralHistory []CollateralHis

type CollateralHis struct {
	ID           int64   `json:"id"`
	CurrencyCode string  `json:"currency_code"`
	Change       float64 `json:"change"`
	Amount       float64 `json:"amount"`
	ReasonCode   string  `json:"reason_code"`
	Date         string  `json:"date"`
}

func (req *CollateralHistory) IsPrivate() bool {
	return true
}

func (req *CollateralHistory) Path() string {
	return "me/getcollateralhistory"
}

func (req *CollateralHistory) Method() string {
	return http.MethodGet
}

func (req *CollateralHistory) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *CollateralHistory) Payload() []byte {
	return nil
}
