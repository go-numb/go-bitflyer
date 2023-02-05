package public

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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

// Optional

func (p *Execution) IsLiquidation() bool {
	if !strings.HasPrefix(p.BuyChildOrderAcceptanceID, "JRF") {
		return true
	}
	if !strings.HasPrefix(p.SellChildOrderAcceptanceID, "JRF") {
		return true
	}

	return false
}

const (
	layout = "20060102-150405"
)

// ToDate changed time from order_id
// Values that appear to be milliseconds below seconds are likely to be random numbers.
func (p *Execution) ToDate() []time.Time {
	ru := []rune(p.BuyChildOrderAcceptanceID)[3:]
	s := string(ru[:15])
	buy, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println(buy)
		return nil
	}

	ru = []rune(p.SellChildOrderAcceptanceID)[3:]
	s = string(ru[:15])
	fmt.Println(s)
	sell, err := time.Parse(layout, s)
	if err != nil {
		return nil
	}

	return []time.Time{buy, sell}
}
