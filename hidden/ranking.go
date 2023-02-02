package hidden

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestRanking struct {
	ContractRegion string `url:"contractRegion,omitempty"`
	Lang           string `url:"lang,omitempty"` // ja
	V              int    `url:"v,omitempty"`
}

type Rankers struct {
	Volume []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"VOLUME"`
	DailyVolume []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"DAILY_VOLUME"`
	YesterdayVolume []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"YESTERDAY_VOLUME"`
	Spot []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"SPOT"`
	Fx []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"FX"`
	DailySpot []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"DAILY_SPOT"`
	DailyFx []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"DAILY_FX"`
	YesterdaySpot []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"YESTERDAY_SPOT"`
	YesterdayFx []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"YESTERDAY_FX"`
	NewMmPoint []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"NEW_MM_POINT"`
	NewMtPoint []struct {
		Nickname       string  `json:"nickname"`
		Volume         float64 `json:"volume"`
		NumberOfTrades int     `json:"number_of_trades"`
		CalculatedDate string  `json:"calculated_date"`
	} `json:"NEW_MT_POINT"`
}

func (p *Client) Ranking(req *RequestRanking) (*Rankers, error) {
	h := http.DefaultClient
	h.Timeout = 20 * time.Second

	u, _ := url.Parse("https://lightning.bitflyer.com/api/trade/ranking")
	u.Path = path.Join("api", "trade", "ranking")
	if req.V != 1 {
		req.V = 1
	}
	v, _ := query.Values(req)
	u.RawQuery = v.Encode()

	fmt.Println(u.String())

	r, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	r.Header.Set("Content-Type", "application/json")

	res, err := h.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 0: 1674993600000
	// 1: 3046415
	// 2: 3049550
	// 3: 3035996
	// 4: 3042148
	// 5: 108.66821058
	// 6: 267.69414682
	// 7: 306.72733209
	// 8: 57.84244411
	// 9: 50.82576647

	rankers := new(Rankers)
	if err := json.NewDecoder(res.Body).Decode(&rankers); err != nil {
		return nil, err
	}

	return rankers, nil
}
