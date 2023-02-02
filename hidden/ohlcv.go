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

type RequestOHLCv struct {
	Symbol string `url:"symbol,omitempty"`
	Period string `url:"period,omitempty"`
	// Count  int    `url:"count,omitempty"`
	// UnixMillisecond
	Before int64 `url:"before,omitempty"`
	// After    int64  `url:"after,omitempty"`
	// Type     string `url:"type,omitempty"`
	// Grouping int    `url:"grouping,omitempty"`
}

type OHLCv struct {
	O float64
	H float64
	L float64
	C float64
	V float64
	// Volume day
	Volume float64

	// for use your target
	// ex) y_target, profit...
	Y float64

	Timestamp time.Time
}

// OHLCv descending-order byte timestamp
// UTC0:00 renew, count max: 720
func (p *Client) OHLCv(req *RequestOHLCv) ([]OHLCv, error) {
	h := http.DefaultClient
	h.Timeout = 3 * time.Second

	u, _ := url.Parse("https://lightchart.bitflyer.com")
	u.Path = path.Join("api", "ohlc")
	v, _ := query.Values(req)
	u.RawQuery = v.Encode()

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

	var data [][]float64
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, fmt.Errorf("response data is nil, has length: %d", len(data))
	}

	var str = make([]OHLCv, len(data))
	for i := 0; i < len(data); i++ {
		var tmp OHLCv
		for j := 0; j < len(data[i]); j++ {
			switch j {
			case 0:
				tmp.Timestamp = time.Unix(int64(data[i][j])/1000, 0)
			case 1:
				tmp.O = data[i][j]
			case 2:
				tmp.H = data[i][j]
			case 3:
				tmp.L = data[i][j]
			case 4:
				tmp.C = data[i][j]
			case 5:
				tmp.V = data[i][j]
			case 6:
				tmp.Volume = data[i][j]
			}
		}
		str[i] = tmp
	}

	return str, nil
}
