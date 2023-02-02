package private

import (
	"net/http"

	"github.com/go-numb/go-bitflyer/types"
	"github.com/google/go-querystring/query"
)

type Executions struct {
	ProductCode            string `url:"product_code,omitempty"`
	Count                  int    `url:"count,omitempty"`
	Before                 int64  `url:"before,omitempty"`
	After                  int64  `url:"after,omitempty"`
	ChildOrderID           string `url:"child_order_id,omitempty"`
	ChildOrderAcceptanceID string `url:"child_order_acceptance_id,omitempty"`
}

type ResponseForExecutions []Execution

type Execution struct {
	ID                     int64   `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	Commission             float64 `json:"commission"`
	ExecDate               string  `json:"exec_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
}

func (req *Executions) IsPrivate() bool {
	return true
}

func (req *Executions) Path() string {
	return "me/getexecutions"
}

func (req *Executions) Method() string {
	return http.MethodGet
}

func (req *Executions) Query() string {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Executions) Payload() []byte {
	return nil
}
