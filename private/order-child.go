package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ChildOrder struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          float64 `json:"price,omitempty"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire,omitempty"`
	TimeInForce    string  `json:"time_in_force,omitempty"`
}

type ResponseForChildOrder struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func (req *ChildOrder) IsPrivate() bool {
	return true
}

func (req *ChildOrder) Path() string {
	return "me/sendchildorder"
}

func (req *ChildOrder) Method() string {
	return http.MethodPost
}

func (req *ChildOrder) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *ChildOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
