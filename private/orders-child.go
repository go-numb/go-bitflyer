package private

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/types"
	"github.com/google/go-querystring/query"
)

type ChildOrders struct {
	ProductCode string `url:"product_code"`
	Count       int    `url:"count,omitempty"`
	Before      int64  `url:"before,omitempty"`
	After       int64  `url:"after,omitempty"`
	// ChildOrderState: ACTIVE, COMPLETED, CANCELED, EXPIRED, REJECTED
	ChildOrderState string `url:"child_order_state,omitempty"`
	ChildOrderID    string `url:"child_order_id,omitempty"`
	ParentOrderID   string `url:"parent_order_id,omitempty"`
}

type ResponseForChildOrders []COrder

type COrder struct {
	ID                     int64   `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	ProductCode            string  `json:"product_code"`
	Side                   string  `json:"side"`
	ChildOrderType         string  `json:"child_order_type"`
	Price                  float64 `json:"price"`
	AveragePrice           float64 `json:"average_price"`
	Size                   float64 `json:"size"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
	TimeInForce            string  `json:"time_in_force"`
}

func (req *ChildOrders) IsPrivate() bool {
	return true
}

func (req *ChildOrders) Path() string {
	return "me/getchildorders"
}

func (req *ChildOrders) Method() string {
	return http.MethodGet
}

func (req *ChildOrders) Query() string {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *ChildOrders) Payload() []byte {
	return nil
}
