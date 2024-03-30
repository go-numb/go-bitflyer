package public

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Fr struct {
	ProductCode string `url:"product_code"`
}

type ResponseForFr struct {
	CurrentFundingRate        float64   `json:"current_funding_rate"`
	NextFundingRateSettledate time.Time `json:"next_funding_rate_settledate"`
}

func (req *Fr) IsPrivate() bool {
	return false
}

func (req *Fr) Path() string {
	return "getfundingrate"
}

func (req *Fr) Method() string {
	return http.MethodGet
}

func (req *Fr) Query() string {
	if req.ProductCode == "" {
		req.ProductCode = "FX_BTC_JPY"
	}

	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Fr) Payload() []byte {
	return nil
}

// custom perse to time.Time
// "next_funding_rate_settledate": "2024-04-15T13:00:00"
// parsing time "2024-03-30T13:00:00" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00"
func (res *ResponseForFr) UnmarshalJSON(b []byte) error {
	type Alias ResponseForFr
	aux := &struct {
		NextFundingRateSettledate string `json:"next_funding_rate_settledate"`
		*Alias
	}{
		Alias: (*Alias)(res),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.NextFundingRateSettledate != "" {
		t, err := time.Parse("2006-01-02T15:04:05", aux.NextFundingRateSettledate)
		if err != nil {
			return err
		}
		res.NextFundingRateSettledate = t
	}

	return nil
}
