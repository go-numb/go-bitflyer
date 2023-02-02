package hidden

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/go-numb/go-bitflyer/types"
)

func TestOHLCv(t *testing.T) {
	hidden := &Client{}
	now := time.Now()
	ohlcv, err := hidden.OHLCv(&RequestOHLCv{
		Symbol: types.FXBTCJPY,
		Period: "m",
		// Count:  500,
		Before: now.UnixMilli(),
		// After: time.Now().Add(-810 * time.Minute).UnixMilli(),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(ohlcv) < 1 {
		t.Fatal("res is nil")
	}

	fmt.Printf("%v to %v\n", ohlcv[0].Timestamp, ohlcv[len(ohlcv)-1].Timestamp)

	res, err := hidden.OHLCv(&RequestOHLCv{
		Symbol: types.FXBTCJPY,
		Period: "m",
		// Count:  500,
		Before: ohlcv[len(ohlcv)-1].Timestamp.Add(-time.Minute).UnixMilli(),
		// After: time.Now().Add(-810 * time.Minute).UnixMilli(),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(res) < 1 {
		t.Fatal("res is nil")
	}

	ohlcv = append(ohlcv, res...)
	fmt.Printf("%v to %v\n", res[0].Timestamp, res[len(res)-1].Timestamp)

	res, err = hidden.OHLCv(&RequestOHLCv{
		Symbol: types.FXBTCJPY,
		Period: "m",
		// Count:  500,
		Before: res[len(res)-1].Timestamp.Add(-time.Minute).UnixMilli(),
		// After: time.Now().Add(-810 * time.Minute).UnixMilli(),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(res) < 1 {
		t.Fatal("res is nil")
	}

	ohlcv = append(ohlcv, res...)

	fmt.Printf("%v to %v\n", res[0].Timestamp, res[len(res)-1].Timestamp)

	fmt.Printf("%v to %v\n", ohlcv[0].Timestamp, ohlcv[len(ohlcv)-1].Timestamp)

	if len(res) < 1 {
		t.Fatal("res is nil")
	}

	// for i := 0; i < len(ohlcv); i++ {
	// 	fmt.Println(ohlcv[i])
	// }

	fmt.Println(len(ohlcv))

	f, _ := os.Create("../_data/chart_minutes.json")
	defer f.Close()
	b, err := json.Marshal(&ohlcv)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write(b); err != nil {
		t.Fatal(err)
	}

}

func TestOHLCvShift(t *testing.T) {
	f, err := os.Open("../_data/chart_minutes.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var ohlcv []OHLCv
	if err := json.NewDecoder(f).Decode(&ohlcv); err != nil {
		t.Fatal(err)
	}

	// // sort ascending-order
	// for i, j := 0, len(ohlcv)-1; i < j; i, j = i+1, j-1 {
	// 	ohlcv[i], ohlcv[j] = ohlcv[j], ohlcv[i]
	// }

	df := dataframe.LoadStructs(
		ohlcv,
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.Float),
		dataframe.WithTypes(
			map[string]series.Type{
				"Timestamp": series.String,
			},
		),
	)
	fmt.Println(df)

	// Shift column: "O"
	v := df.Col("O").Float()
	v = append([]float64{math.NaN(), math.NaN(), math.NaN()}, v[:len(v)-3]...)
	df = df.Mutate(
		series.New(v, series.Float, "O"),
	)
	fmt.Println(df)
}

func TestRanking(t *testing.T) {
	hidden := &Client{}
	res, err := hidden.Ranking(&RequestRanking{
		ContractRegion: "JP",
	})
	if err != nil {
		t.Fatal(err)
	}

	f, _ := os.Create("../_data/ranking.json")
	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write(b); err != nil {
		t.Fatal(err)
	}
	f.Close()
}
