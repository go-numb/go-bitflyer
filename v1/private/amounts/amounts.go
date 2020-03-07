package amounts

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct{}

type Response []Account

type Account struct {
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Amount       float64            `json:"amount"`
}

const (
	APIPath string = "/v1/me/getcollateralhistory"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
