# go-bitflyer
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)



Bitflyer exchange API version1, renew at 2023/02.

## Description

go-bitflyer is a go client library for [Bitflyer lightning API](https://lightning.bitflyer.com/docs).

**Supported**
- [x] All Public API
- [x] All Private API
- [x] GUI's Hidden API
- [x] Websocket


# Usage
| Client   | Functions                                                                                     |
| -------- | --------------------------------------------------------------------------------------------- |
| Public   | Status(&Request{}) (res, apiLimit, err)                                                       |
| Public   | Helth(&Request{}) (res, apiLimit, err)                                                        |
| Public   | Ticker(&Request{}) (res, apiLimit, err)                                                       |
| Public   | Market(&Request{}) (res, apiLimit, err)                                                       |
| Public   | Board(&Request{}) (res, apiLimit, err)                                                        |
| Public   | Executions(&Request{})                                                                        |
| Public   | Chat(&Request{}) (res, apiLimit, err)                                                         |
| Private  | Permissions(&Request{}) (res, apiLimit, err)                                                  |
| Private  | Balance(&Request{}) (res, apiLimit, err)                                                      |
| Private  | Collateral(&Request{}) (res, apiLimit, err)                                                   |
| Private  | CollateralAccounts(&Request{}) (res, apiLimit, err)                                           |
| Private  | Addresses(&Request{}) (res, apiLimit, err)                                                    |
| Private  | Coinins(&Request{}) (res, apiLimit, err)                                                      |
| Private  | Coinouts(&Request{}) (res, apiLimit, err)                                                     |
| Private  | Banks(&Request{}) (res, apiLimit, err)                                                        |
| Private  | Deposits(&Request{}) (res, apiLimit, err)                                                     |
| Private  | Withdrawals(&Request{}) (res, apiLimit, err)                                                  |
| Private  | ChildOrder(&Request{}) (res, apiLimit, err)                                                   |
| Private  | CancelChildOrder(&Request{}) (res, apiLimit, err)                                             |
| Private  | ParentOrder(&Request{}) (res, apiLimit, err)                                                  |
| Private  | CancelParentOrder(&Request{}) (res, apiLimit, err)                                            |
| Private  | Cancel(&Request{}) (res, apiLimit, err)                                                       |
| Private  | ChildOrders(&Request{}) (res, apiLimit, err)                                                  |
| Private  | DetailParentOrder(&Request{}) (res, apiLimit, err)                                            |
| Private  | MyExecutions(&Request{}) (res, apiLimit, err)                                                 |
| Private  | BalanceHistory(&Request{}) (res, apiLimit, err)                                               |
| Private  | Positions(&Request{}) (res, apiLimit, err)                                                    |
| Private  | CollateralHistory(&Request{}) (res, apiLimit, err)                                            |
| Private  | Commission(&Request{}) (res, apiLimit, err)                                                   |
| Hiddn    | OHLCv(&Request{}) (res, err)                                                                  |
| Hiddn    | Ranking(&Request{}) (res, err)                                                                |
| Realtime | Connect(&auth, []string{"channels"...}, []string{"product_codes..."...}, chan Response) error |


```go HTTP API

package main

func main() {
    client := bitflyer.New(os.Getenv("BF_KEY"), os.Getenv("BF_SECRET"))

    results, managedApiLimit, err := client.Executions(&public.Executions{
        ProductCode: "BTC_JPY",
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range *res {
        fmt.Println(v)
        // -> {2430391013 BUY 2.989057e+06 0.02 2025-01-01T08:47:20.577 JRF20250101-084720-050159 JRF20250101-084720-042209}...
    }
    fmt.Println(manage.Remain)
    // API Limit remain and more
    // -> 489

	// For Ohlcv and Ranking
	hclient := hidden.New()
	ohlcv, err := hclient.OHLCv(&hidden.RequestOHLCv{
        Symbo;: "BTC_JPY",
		Period: "m",
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range ohlcv {
        fmt.Println(v)
        // -> C, H, L, O, Timestamp, V 
		// 0: 3003659 3004234 3003143 3003503 2023-02-01T21:01:00+09:00 3.131921
    }
}


```


```go Realtime API with websocket

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-numb/go-bitflyer"
	"github.com/go-numb/go-bitflyer/public"
	"github.com/go-numb/go-bitflyer/realtime"
	"github.com/go-numb/go-bitflyer/types"
)

func main() {
	client := bitflyer.New("xxxxxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxx=")

    // cancel goroutine from this function
	ctx, cancel := context.WithCancel(context.Background())
	ws := realtime.New(ctx)
	defer ws.Close()

    // recive data and notification form websocket
	recived := make(chan realtime.Response, 3)
	go ws.Connect(
        // Connect to Private, if struct *Auth. 
		client.Auth, // or nil, when not subscribe to private channel

        // input channel and product code, create request paramater in function
		[]string{
			"executions",
			"ticker",
			"board",
		},
		[]string{
			types.BTCJPY,
		},
		recived)

	go Reciver(ctx, recived)

    sig := make(chan os.Signal)
    <-sig
	cancel()
	time.Sleep(time.Second)
	close(recived)

}

func Reciver(ctx context.Context, ch chan realtime.Response) {
	defer fmt.Println("func reciver was done")
L:
	for {
		select {
		case <-ctx.Done():
			break L
		case v := <-ch:
			switch v.Types {
			case realtime.TickerN:
				fmt.Print(v.ProductCode)
				fmt.Printf(" ticker comming, %.0f\n", v.Ticker.Ltp)
			case realtime.ExecutionsN:
				fmt.Print(v.ProductCode)
				fmt.Println(" executions comming, id: ", v.Executions[0].ID)
			case realtime.BoardSnapN:
				fmt.Print(v.ProductCode)
				fmt.Println(" board snapshot comming, ", v.Board.MidPrice)
			case realtime.BoardN:
				fmt.Print(v.ProductCode)
				fmt.Println(" board update comming, ", len(v.Board.Asks), len(v.Board.Bids))

			case realtime.ChildOrdersN:
				fmt.Println("child order comming, ", v.ChildOrders[0].ChildOrderID)
			case realtime.ParentOrdersN:
				fmt.Println("parent order comming, ", v.ParentOrders[0].ParentOrderID)

			default:
				fmt.Printf("error or undefind: %#v", v)
			}

		}
	}
}


```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitflyer/blob/master/LICENSE)