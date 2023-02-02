package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type DetailParentOrder struct {
	ParentOrderID           string `url:"parent_order_id,omitempty"`
	ParentOrderAcceptanceID string `url:"parent_order_acceptance_id,omitempty"`
}

type ResponseForDetailParentOrder struct {
	ID                      int64   `json:"id"`
	ParentOrderID           string  `json:"parent_order_id"`
	OrderMethod             string  `json:"order_method"`
	ExpireDate              string  `json:"expire_date"`
	TimeInForce             string  `json:"time_in_force"`
	Parameters              []Param `json:"parameters"`
	ParentOrderAcceptanceID string  `json:"parent_order_acceptance_id"`
}

func (req *DetailParentOrder) IsPrivate() bool {
	return true
}

func (req *DetailParentOrder) Path() string {
	return "me/getparentorder"
}

func (req *DetailParentOrder) Method() string {
	return http.MethodGet
}

func (req *DetailParentOrder) Query() string {
	if req.ParentOrderID != "" && req.ParentOrderAcceptanceID != "" {
		req.ParentOrderAcceptanceID = ""
	}
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *DetailParentOrder) Payload() []byte {
	return nil
}
