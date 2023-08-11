package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"

	"github.com/rluisr/go-bitflyer/types"
)

type Cancel struct {
	ProductCode string `json:"product_code"`
}

type ResponseForCancel struct{}

func (req *Cancel) IsPrivate() bool {
	return true
}

func (req *Cancel) Path() string {
	return "me/cancelallchildorders"
}

func (req *Cancel) Method() string {
	return http.MethodPost
}

func (req *Cancel) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Cancel) Payload() []byte {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
