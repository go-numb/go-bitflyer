package bitflyer

import (
	"math"
)

const (
	DECIMALP = 1e8
	DECIMALM = 1e-8
)

// Round for order size
// 最小注文単位以下の枚数
func Round(f float64) float64 {
	return math.Round(f*DECIMALP) * DECIMALM
}
