package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type LeverageC struct {
}

type ResponseForLeverageC struct {
	CurrentMax       float64 `json:"current_max"`
	CurrentStartdate string  `json:"current_startdate"`
	NextMax          float64 `json:"next_max"`
	NextStartdate    string  `json:"next_startdate"`
}

func (req *LeverageC) IsPrivate() bool {
	return false
}

func (req *LeverageC) Path() string {
	return "getcorporateleverage"
}

func (req *LeverageC) Method() string {
	return http.MethodGet
}

func (req *LeverageC) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *LeverageC) Payload() []byte {
	return nil
}
