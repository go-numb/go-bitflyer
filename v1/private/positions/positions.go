package positions

import (
	"math"
	"net/http"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode types.ProductCode `json:"product_code" url:"product_code"`
}

type Position struct {
	ProductCode         types.ProductCode `json:"product_code"`
	Side                string            `json:"side"`
	Price               float64           `json:"price"`
	Size                float64           `json:"size"`
	Commission          float64           `json:"commission"`
	SwapPointAccumulate float64           `json:"swap_point_accumulate"`
	RequireCollateral   float64           `json:"require_collateral"`
	OpenDate            time.BitflyerTime `json:"open_date"`
	Leverage            float64           `json:"leverage"`
	Pnl                 float64           `json:"pnl"`
	Sfd                 float64           `json:"sfd"`
}

type Response []Position

const (
	APIPath string = "/v1/me/getpositions"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}

/*
	# Positions managed with tanh()
*/

// T is Positions struct
type T struct {
	Min   float64
	Size  float64
	Limit float64
}

// NewT is new Positions struct
func NewT(min, limit float64) *T {
	return &T{
		Min:   min,
		Limit: limit,
	}
}

// Set is sets size
func (p *T) Set(size float64) {
	p.Size = size
}

// Lot is Size for order lot
// p.Sizeに対応する解消符号でsizeが返ってくる
func (p *T) Lot(side int, tension float64) (bool, float64) {
	if p.isFull(side) {
		return true, 0
	}

	var adjust = 0.001

	lot := math.RoundToEven(math.Abs(p.Limit*p.bias(tension))/adjust) * adjust
	if lot < p.Min {
		return false, p.Min
	}
	return false, math.Min(p.Limit, lot)
}

// bias is positions bias
func (p *T) bias(tension float64) float64 {
	return math.Tanh(tension * p.Size / p.Limit)
}

// isFull is checks Limit&Size
// 売り方向要望を受け、同方向建玉過多ならばisFull
func (p *T) isFull(side int) bool {
	if 0 < side && p.Limit < p.Size {
		return true
	} else if side < 0 && p.Size < -p.Limit {
		return true
	}

	return false
}
