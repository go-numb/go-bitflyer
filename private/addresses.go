package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Addresses struct {
}

type ResponseForAddresses []Address

type Address struct {
	Type         string `json:"type"`
	CurrencyCode string `json:"currency_code"`
	Address      string `json:"address"`
}

func (req *Addresses) IsPrivate() bool {
	return true
}

func (req *Addresses) Path() string {
	return "me/getaddresses"
}

func (req *Addresses) Method() string {
	return http.MethodGet
}

func (req *Addresses) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Addresses) Payload() []byte {
	return nil
}
