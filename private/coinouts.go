package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Coinouts struct {
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

type ResponseForCoinouts []Coinout

type Coinout struct {
	ID            int     `json:"id"`
	OrderID       string  `json:"order_id"`
	CurrencyCode  string  `json:"currency_code"`
	Amount        float64 `json:"amount"`
	Address       string  `json:"address"`
	TxHash        string  `json:"tx_hash"`
	Fee           float64 `json:"fee"`
	AdditionalFee float64 `json:"additional_fee"`
	Status        string  `json:"status"`
	EventDate     string  `json:"event_date"`
}

func (req *Coinouts) IsPrivate() bool {
	return true
}

func (req *Coinouts) Path() string {
	return "me/getcoinouts"
}

func (req *Coinouts) Method() string {
	return http.MethodGet
}

func (req *Coinouts) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Coinouts) Payload() []byte {
	return nil
}
