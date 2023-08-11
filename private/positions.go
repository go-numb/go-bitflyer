package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/rluisr/go-bitflyer/types"
)

type Positions struct {
	ProductCode string `url:"product_code,omitempty"`
}

type ResponseForPositions []Position

type Position struct {
	ProductCode         string  `json:"product_code"`
	Side                string  `json:"side"`
	Price               float64 `json:"price"`
	Size                float64 `json:"size"`
	Commission          float64 `json:"commission"`
	SwapPointAccumulate float64 `json:"swap_point_accumulate"`
	RequireCollateral   float64 `json:"require_collateral"`
	OpenDate            string  `json:"open_date"`
	Leverage            float64 `json:"leverage"`
	Pnl                 float64 `json:"pnl"`
	Sfd                 float64 `json:"sfd"`
}

func (req *Positions) IsPrivate() bool {
	return true
}

func (req *Positions) Path() string {
	return "me/getpositions"
}

func (req *Positions) Method() string {
	return http.MethodGet
}

func (req *Positions) Query() string {
	if req.ProductCode == "" {
		req.ProductCode = types.BTCJPY
	}
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Positions) Payload() []byte {
	return nil
}
