package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ParentOrder struct {
	OrderMethod    string  `json:"order_method,omitempty"`
	MinuteToExpire int     `json:"minute_to_expire,omitempty"`
	TimeInForce    string  `json:"time_in_force,omitempty"`
	Parameters     []Param `json:"parameters"`
}

type Param struct {
	ProductCode   string  `json:"product_code"`
	ConditionType string  `json:"condition_type"`
	Side          string  `json:"side"`
	Price         float64 `json:"price,omitempty"`
	Size          float64 `json:"size"`
	TriggerPrice  float64 `json:"trigger_price,omitempty"`
	Offset        float64 `json:"offset,omitempty"`
}

type ResponseForParentOrder struct {
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id"`
}

func (req *ParentOrder) IsPrivate() bool {
	return true
}

func (req *ParentOrder) Path() string {
	return "me/sendparentorder"
}

func (req *ParentOrder) Method() string {
	return http.MethodPost
}

func (req *ParentOrder) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *ParentOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
