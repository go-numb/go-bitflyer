package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Executions struct {
	ProductCode string `url:"product_code,omitempty"`
	// MAX 過去 31 日分
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

type ResponseForExecutions []Execution

type Execution struct {
	ID                         int64   `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

func (req *Executions) IsPrivate() bool {
	return false
}

func (req *Executions) Path() string {
	return "executions"
}

func (req *Executions) Method() string {
	return http.MethodGet
}

func (req *Executions) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Executions) Payload() []byte {
	return nil
}
