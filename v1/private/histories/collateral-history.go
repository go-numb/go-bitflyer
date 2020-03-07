package histories

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-numb/go-bitflyer/v1/time"
	"github.com/go-numb/go-bitflyer/v1/types"
)

type Request struct {
	Pagination types.Pagination `json:",inline"`
}

type Response []History

type History struct {
	ID           int               `json:"id"`
	CurrencyCode string            `json:"currency_code"`
	Change       float64           `json:"change"`
	Amount       float64           `json:"amount"`
	ReasonCode   string            `json:"reason_code"`
	Date         time.BitflyerTime `json:"date"`
}

const (
	APIPath string = "/v1/me/getcollateralhistory"
)

type SFDFactors struct {
	successCount, failCount int
	success, fail           float64
}

func (p *Response) SFDFactor() *SFDFactors {
	s := new(SFDFactors)

	for _, v := range *p {
		if 0 < v.Change {
			s.successCount++
			s.success += v.Change
		} else if v.Change < 0 {
			s.failCount++
			s.fail += v.Change
		}
	}
	return s
}

func (s *SFDFactors) Culc() (countF, sfdF float64) {
	return math.Max(0, float64(s.successCount)/float64(s.failCount)), math.Max(0, s.success/math.Abs(s.fail))
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	// values, _ := query.Values(req)
	var q []string
	if !reflect.DeepEqual(req.Pagination, types.Pagination{}) {
		if req.Pagination.Count != 0 {
			q = append(q, fmt.Sprintf("count=%d", req.Pagination.Count))
		}
		if req.Pagination.Before != 0 {
			q = append(q, fmt.Sprintf("before=%d", req.Pagination.Before))
		}
		if req.Pagination.After != 0 {
			q = append(q, fmt.Sprintf("after=%d", req.Pagination.After))
		}
	}

	return strings.Join(q, "&")
}

func (req *Request) Payload() []byte {
	return nil
}
