package v1

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	// MARKET is 成行き
	MARKET = "MARKET"
	// LIMIT is 指値
	LIMIT = "LIMIT"
	// BUY is 買い注文
	BUY = "BUY"
	// SELL is 売り注文
	SELL = "SELL"

	// Type TimeInForce
	IOC = "IOC"
	FOK = "FOK"

	APIREMAIN  = 500
	TIMELAYOUT = "20060102.150405.999999999"
)

func ToType(isMarket bool) string {
	if isMarket {
		return MARKET
	}
	return LIMIT
}

func ToSide(side int) string {
	if side == 1 {
		return BUY
	}
	return SELL
}

func ToPrice(price float64) float64 {
	return math.Abs(math.RoundToEven(price))
}

func ToSize(size float64) float64 {
	size = math.Abs(size)
	if size < 0.01 {
		return 0.01
	}
	return math.Floor(size*10000) / 10000
}

func ToTimeByOrderID(s string) (time.Time, error) {
	// Prefix を削除しつつ，format string for parse
	s = strings.Replace(s[3:], "-", ".", -1)

	return time.Parse(TIMELAYOUT, s)
}

type API struct {
	url string
}

func NewAPI(c *Client, apiPath string) *API {
	return &API{
		url: c.APIHost() + apiPath,
	}
}

func (api *API) ToURL() (*url.URL, error) {
	u, err := url.ParseRequestURI(api.url)
	if err != nil {
		return nil, errors.Wrapf(err, "can not parse url: %s", api.url)
	}
	return u, nil
}

// 基本的には5分毎リセット
type APIHeaders struct {
	Public  Limit
	Private Limit
}

func (p *APIHeaders) IsCache(h http.Header) bool {
	isCache := h.Get("Pragma")
	if isCache != "no-cache" {
		// on キャッシュ
		return true
	}
	// does not キャッシュ
	return false
}

type Limit struct {
	Period int       // Period is リセットまでの秒数
	Remain int       // Remain is 残Requests
	Reset  time.Time // Reset Remainの詳細時間(sec未満なし)
}

func NewLimit(isPrivate bool) *Limit {
	if isPrivate {
		return &Limit{
			Period: 0,
			Remain: APIREMAIN,
			Reset:  time.Now().Add(5 * time.Minute),
		}
	}

	return &Limit{
		Period: 0,
		Remain: APIREMAIN,
		Reset:  time.Now().Add(5 * time.Minute),
	}
}

// FromHeader X-xxxからLimitを取得
func (p *Limit) FromHeader(h http.Header) {
	period := h.Get("X-Ratelimit-Period") // リセットまでの残秒数
	if period != "" {
		p.Period, _ = strconv.Atoi(period)
	}
	remain := h.Get("X-Ratelimit-Remaining") // 残回数
	if remain != "" {
		p.Remain, _ = strconv.Atoi(remain)
	}
	t := h.Get("X-Ratelimit-Reset") // リセットUTC時間(sec未満なし)
	if t != "" {
		reset, _ := strconv.ParseInt(t, 10, 64)
		p.toTime(reset)
	}
}

func (p *Limit) Check() error {
	if p.Remain <= 0 {
		if time.Now().After(p.Reset) { // APIRESET時間を過ぎていたらRemainを補充
			p.Remain = APIREMAIN
		}
		return fmt.Errorf("api limit, has API Limit Remain:%d, Reset time: %s(%s)",
			p.Remain,
			p.Reset.Format("15:04:05"),
			time.Now().Format("15:04:05"))
	}
	return nil
}

// int64 to time.Time
func (p *Limit) toTime(t int64) {
	p.Reset = time.Unix(t, 10)
}
