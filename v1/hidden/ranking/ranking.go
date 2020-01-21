package ranking

import (
	"time"

	"github.com/buger/jsonparser"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Ranker struct {
	CalculatedDate     string    `json:"calculated_date,omitempty"`
	Nickname           string    `json:"nickname,omitempty"`
	NumberOfTrades     int       `json:"number_of_trades,omitempty"`
	Volume             float64   `json:"volume,omitempty"`
	DiffNumberOfTrades int       `json:"diff_number_of_trades,omitempty"`
	DiffVolume         float64   `json:"diff_volume,omitempty"`
	DiffOneShot        float64   `json:"diff_one_shot,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}

const RANKURL = "https://lightning.bitflyer.com/api/trade/ranking?contractRegion=JP"

// Get open ranking informations, day & week
// key is VOLUME or DAYLY ...
func Get(key string) ([]Ranker, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(RANKURL)

	res := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}

	data, _, _, err := jsonparser.Get(res.Body(), key)
	if err != nil {
		return nil, err
	}
	var rank []Ranker
	json.Unmarshal(data, &rank)

	return rank, nil
}
