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
// p.Sizeに対応するsizeが返ってくる
func (p *T) Lot(side int, tension float64) (bool, float64) {

	bias := p.bias(tension)
	size := p.Limit * bias
	if p.isFull(side, size) {
		return true, 0
	}

	size = p.checkSame(side, size)

	return false, math.Max(p.Min, math.Abs(size))
}

// bias is positions bias
func (p *T) bias(tension float64) float64 {
	return math.Tanh(tension * p.Size / p.Limit)
}

// isFull is checks Limit&Size
// 売り方向要望を受け、同方向建玉過多ならばisFull
func (p *T) isFull(side int, size float64) bool {
	if 0 < side { // 新規買い注文
		if p.Limit < p.Size {
			return true
		}
		if p.Limit < p.Size+math.Abs(size) {
			return true
		}
		if p.Limit < p.Size+p.Min {
			return true
		}

	} else if side < 0 { // 新規売り注文
		if p.Size < -p.Limit {
			return true
		}
		if p.Size-math.Abs(size) < -p.Limit {
			return true
		}
		if p.Size-math.Abs(p.Min) < -p.Limit {
			return true
		}

	}

	return false
}

func (p *T) checkSame(side int, size float64) float64 {
	if p.Min < math.Abs(size) { // 注文多重化の際、買い建玉に買い注文過多を避ける目的
		if 0 < side && 0 < p.Size {
			size = p.Min
		} else if side < 0 && p.Size < 0 {
			size = p.Min
		}
	}
	return size
}
