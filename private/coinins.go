package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Coinins struct {
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

type ResponseForCoinins []Coinin

type Coinin struct {
	ID           int     `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Address      string  `json:"address"`
	TxHash       string  `json:"tx_hash"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (req *Coinins) IsPrivate() bool {
	return true
}

func (req *Coinins) Path() string {
	return "me/getcoinins"
}

func (req *Coinins) Method() string {
	return http.MethodGet
}

func (req *Coinins) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Coinins) Payload() []byte {
	return nil
}
