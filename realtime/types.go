package realtime

import (
	"github.com/go-numb/go-bitflyer/public"
)

const (
	JSONRPCV    = "2.0"
	AUTH        = "auth"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"

	CHANNEL = "channel"
)

type Types int

const (
	AllN Types = iota
	TickerN
	ExecutionsN
	BoardN
	BoardSnapN
	ChildOrdersN
	ParentOrdersN
	UndefinedN
	ErrorN
)

const (
	ALL          = "ALL"
	Ticker       = "lightning_ticker_"
	Executions   = "lightning_executions_"
	Board        = "lightning_board_"
	BoardSnap    = "lightning_board_snapshot_"
	ChildOrders  = "child_order_events"
	ParentOrders = "parent_order_events"
	Error        = "error"
	Undefined    = "undefined"
)

func (p Types) String() string {
	switch p {
	case AllN:
		return ALL
	case TickerN:
		return Ticker
	case ExecutionsN:
		return Executions
	case BoardN:
		return Board
	case BoardSnapN:
		return BoardSnap
	case ChildOrdersN:
		return ChildOrders
	case ParentOrdersN:
		return ParentOrders
	case ErrorN:
		return Error
	}
	return Undefined
}

type Request struct {
	Jsonrpc string                 `json:"jsonrpc,omitempty"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      int                    `json:"id,omitempty"`
}

func (p *Response) _set(err error) Response {
	p.Types = ErrorN
	p.Results = err
	return *p
}

type Response struct {
	Types       Types
	ProductCode string

	Board      public.ResponseForBoard
	Ticker     public.ResponseForTicker
	Executions []public.Execution

	ChildOrders  []ChildOrder
	ParentOrders []ParentOrder

	Results error
}

type ChildOrder struct {
	ProductCode            string  `json:"product_code"`
	ChildOrderID           string  `json:"child_order_id"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	EventDate              string  `json:"event_date"`
	EventType              string  `json:"event_type"`
	ExecID                 int     `json:"exec_id,omitempty"`
	ChildOrderType         string  `json:"child_order_type,omitempty"`
	Side                   string  `json:"side,omitempty"`
	Price                  int     `json:"price"`
	Size                   float64 `json:"size"`
	ExpireDate             string  `json:"expire_date,omitempty"`
}

type ParentOrder struct {
	ProductCode             string  `json:"product_code"`
	ParentOrderID           string  `json:"parent_order_id"`
	ParentOrderAcceptanceID string  `json:"parent_order_acceptance_id"`
	EventDate               string  `json:"event_date"`
	EventType               string  `json:"event_type"`
	ParameterIndex          int     `json:"parameter_index"`
	ChildOrderType          string  `json:"child_order_type,omitempty"`
	Side                    string  `json:"side,omitempty"`
	Price                   int     `json:"price"`
	Size                    float64 `json:"size,omitempty"`
	ExpireDate              string  `json:"expire_date,omitempty"`
	ChildOrderAcceptanceID  string  `json:"child_order_acceptance_id"`
}
