package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type ParentOrders struct {
	ProductCode int   `url:"product_code,omitempty"`
	Count       int   `url:"count,omitempty"`
	Before      int64 `url:"before,omitempty"`
	// ChildOrderState: ACTIVE, COMPLETED, CANCELED, EXPIRED, REJECTED
	ParentOrderState string `url:"parent_order_state,omitempty"`
}

type ResponseForParentOrders []POrder

type POrder struct {
	ID                      int64   `json:"id"`
	ParentOrderID           string  `json:"parent_order_id"`
	ProductCode             string  `json:"product_code"`
	Side                    string  `json:"side"`
	ParentOrderType         string  `json:"parent_order_type"`
	Price                   float64 `json:"price"`
	AveragePrice            float64 `json:"average_price"`
	Size                    float64 `json:"size"`
	ParentOrderState        string  `json:"parent_order_state"`
	ExpireDate              string  `json:"expire_date"`
	ParentOrderDate         string  `json:"parent_order_date"`
	ParentOrderAcceptanceID string  `json:"parent_order_acceptance_id"`
	OutstandingSize         float64 `json:"outstanding_size"`
	CancelSize              float64 `json:"cancel_size"`
	ExecutedSize            float64 `json:"executed_size"`
	TotalCommission         float64 `json:"total_commission"`
}

func (req *ParentOrders) IsPrivate() bool {
	return true
}

func (req *ParentOrders) Path() string {
	return "me/getparentorders"
}

func (req *ParentOrders) Method() string {
	return http.MethodGet
}

func (req *ParentOrders) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *ParentOrders) Payload() []byte {
	return nil
}
