package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/rluisr/go-bitflyer/types"
)

type CancelChildOrder struct {
	ProductCode            string `json:"product_code"`
	ChildOrderID           string `json:"child_order_id,omitempty"`
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id,omitempty"`
}

type ResponseForCancelChildOrder struct{}

func (req *CancelChildOrder) IsPrivate() bool {
	return true
}

func (req *CancelChildOrder) Path() string {
	return "me/cancelchildorder"
}

func (req *CancelChildOrder) Method() string {
	return http.MethodPost
}

func (req *CancelChildOrder) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *CancelChildOrder) Payload() []byte {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	if req.ChildOrderID != "" && req.ChildOrderAcceptanceID != "" {
		req.ChildOrderAcceptanceID = ""
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
