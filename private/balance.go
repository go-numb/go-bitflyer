package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Balance struct {
}

type ResponseForBalance []Currency

type Currency struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available,omitempty"`
}

func (req *Balance) IsPrivate() bool {
	return true
}

func (req *Balance) Path() string {
	return "me/getbalance"
}

func (req *Balance) Method() string {
	return http.MethodGet
}

func (req *Balance) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Balance) Payload() []byte {
	return nil
}
