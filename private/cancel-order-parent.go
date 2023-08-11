package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/rluisr/go-bitflyer/types"
)

type CancelParentOrder struct {
	ProductCode             string `json:"product_code"`
	ParentOrderID           string `json:"parent_order_id,omitempty"`
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id,omitempty"`
}

type ResponseForCancelParentOrder struct{}

func (req *CancelParentOrder) IsPrivate() bool {
	return true
}

func (req *CancelParentOrder) Path() string {
	return "me/cancelparentorder"
}

func (req *CancelParentOrder) Method() string {
	return http.MethodPost
}

func (req *CancelParentOrder) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *CancelParentOrder) Payload() []byte {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	if req.ParentOrderID != "" && req.ParentOrderAcceptanceID != "" {
		req.ParentOrderAcceptanceID = ""
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
