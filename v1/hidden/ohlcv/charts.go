package ohlcv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	OHLCv = "https://lightchart.bitflyer.com/api/ohlc?symbol=%s&period=m&after=%d&before=%d"
)

// Get is unsupported by Bitflyer
// ※ Reset length:0 at 21:00 JTC
func Get(minutes int, productCode string) (length int, ohclv [][]float64, err error) {
	url := fmt.Sprintf(
		OHLCv,
		productCode,
		time.Now().UTC().Add(-time.Duration(minutes)*time.Minute).UnixNano()/1000000,
		time.Now().UTC().UnixNano()/1000000)
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	if err != nil {
		return 0, nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	var ff [][]float64
	json.NewDecoder(res.Body).Decode(&ff)

	fmt.Printf("set ohlcv, not enough at bitflyer, lenght is %d\n", len(ff))

	ln := len(ff)
	if ln < minutes {
		url = fmt.Sprintf("https://api.cryptowat.ch/markets/bitflyer/fxbtcjpy/ohlc?periods=60&after=%d&before=%d",
			time.Now().UTC().Add(-180*time.Minute).Unix(),
			time.Now().UTC().Unix())
		req, err = http.NewRequest(
			http.MethodGet,
			url,
			nil)
		if err != nil {
			return 0, nil, err
		}
		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return 0, nil, err
		}
		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(&ff)
		fmt.Println("use cryptowatch ohlcv")
	}

	return ln, ff, nil
}
