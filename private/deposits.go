package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Deposits struct {
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

type ResponseForDeposits []Deposit

type Deposit struct {
	ID           int     `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (req *Deposits) IsPrivate() bool {
	return true
}

func (req *Deposits) Path() string {
	return "me/getdeposits"
}

func (req *Deposits) Method() string {
	return http.MethodGet
}

func (req *Deposits) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Deposits) Payload() []byte {
	return nil
}
