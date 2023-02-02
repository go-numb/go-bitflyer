package private

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/types"
	"github.com/google/go-querystring/query"
)

type Commission struct {
	ProductCode string `url:"product_code"`
}

type ResponseForCommission struct {
	CommissionRate float64 `json:"commission_rate"`
}

func (req *Commission) IsPrivate() bool {
	return true
}

func (req *Commission) Path() string {
	return "me/gettradingcommission"
}

func (req *Commission) Method() string {
	return http.MethodGet
}

func (req *Commission) Query() string {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Commission) Payload() []byte {
	return nil
}
