package types

type CurrencyCode string

type ProductCode string

const (
	// type ProductCode
	BTCJPY   ProductCode = "BTC_JPY"
	FXBTCJPY ProductCode = "FX_BTC_JPY"
	ETHJPY   ProductCode = "ETH_JPY"
	ETHBTC   ProductCode = "ETH_BTC"
	// BTCFUTURE1 = "BTCJPY27SEP2019"
	// BTCFUTURE2 = "BTCJPY30AUG2019"
	// BTCFUTUREx = "BTCJPY06SEP2019"
	UNDEFINED ProductCode = "undefined"
)

type Pagination struct {
	Count  int `json:"count,omitempty" url:"count,omitempty"`
	Before int `json:"before,omitempty" url:"before,omitempty"`
	After  int `json:"after,omitempty" url:"after,omitempty"`
}
