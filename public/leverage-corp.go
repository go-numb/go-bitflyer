package public

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type LeverageC struct {
}

type ResponseForLeverageC struct {
	CurrentMax       float64   `json:"current_max"`
	CurrentStartdate time.Time `json:"current_startdate"`
	NextMax          float64   `json:"next_max"`
	NextStartdate    time.Time `json:"next_startdate"`
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

func (res *ResponseForLeverageC) UnmarshalJSON(b []byte) error {
	type Alias ResponseForLeverageC
	aux := &struct {
		CurrentStartdate string `json:"current_startdate"`
		NextStartdate    string `json:"next_startdate"`
		*Alias
	}{
		Alias: (*Alias)(res),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.CurrentStartdate != "" {
		t, err := time.Parse("2006-01-02T15:04:05", aux.CurrentStartdate)
		if err != nil {
			return err
		}
		res.CurrentStartdate = t
	}
	if aux.NextStartdate != "" {
		t, err := time.Parse("2006-01-02T15:04:05", aux.NextStartdate)
		if err != nil {
			return err
		}
		res.NextStartdate = t
	}

	return nil
}
