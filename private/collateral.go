package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Collateral struct {
}

type ResponseForCollateral struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
	MarginCallAmount  float64 `json:"margin_call_amount"`
	MarginCallDueDate string  `json:"margin_call_due_date"`
}

func (req *Collateral) IsPrivate() bool {
	return true
}

func (req *Collateral) Path() string {
	return "me/getcollateral"
}

func (req *Collateral) Method() string {
	return http.MethodGet
}

func (req *Collateral) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Collateral) Payload() []byte {
	return nil
}
