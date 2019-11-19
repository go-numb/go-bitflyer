# go-bitflyer

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-bitflyer is wrapper for Crypto Trading [bitFlyer Lightning API](https://lightning.bitflyer.com/docs), with Golang.


# Fork & Tribute
[github@kkohtaka](https://github.com/kkohtaka/go-bitflyer)

## Modifications
- bitflyer.com
- time.UTC()
- Akamai, and user headers
- types
- times.Bitflyer
- Cancel by id
- Order's special
- API Limit from headers
- API data cached
- Websocket for private(child/parent orders)

## Usage

```golang
package main

import (
  "log"

  "github.com/go-numb/go-bitflyer/auth"
  "github.com/go-numb/go-bitflyer/v1"
  "github.com/go-numb/go-bitflyer/v1/public/markets"
  "github.com/go-numb/go-bitflyer/v1/private/permissions"
)

func main() {
  client := v1.NewClient(&v1.ClientOpts{
    AuthConfig: &auth.AuthConfig{
      APIKey:    "<api_key>",
      APISecret: "<api_secretkey>",
    },
  })

  // return this Struct, Raw response, error 
  s, res, err := client.Permissions(&permissions.Request{})
  if err != nil {
    log.Fatalln(err)
  } else {
    log.Println(resp)
  }

  s, res, err = client.Markets(&markets.Request{})
  if err != nil {
    log.Fatalln(err)
  } else {
    log.Println(resp)
  }
}

```


# bitflyer API realtime json-rpc
```golang
import	"github.com/go-numb/go-bitflyer/v1/jsonrpc"

func main() {
  dataStruct := struct{}

  // Websocket JsonRPC
	ch := make(chan jsonrpc.Response)
	channels := []string{
		"lightning_board_FX_BTC_JPY",
		"lightning_ticker_FX_BTC_JPY",
    "lightning_executions_FX_BTC_JPY",
  }
  channelsPrivate := []string{
    // spot tickerを購読する理由
    // - 受信待機時間が長くなる場合、設定したタイムアウト時間内に一度も受信しない可能性
    // - bitflyer JSON-RPCにはpingによるpongがない
    // - 受信しない場合タイムアウト設定を設けている
    // 上記理由により、ticker配信にて受信Eventを行っている
    "lightning_ticker_BTC_JPY",
    "child_order_events",
    "parent_order_events",
  }
  go jsonrpc.Get(channels, ch)
  go jsonrpc.GetPrivate(<KEY>, <SECRET>, channelsPrivate, ch)

	eg.Go(func() error {
		for {
			select {
			case v := <-ch: // read websocket
				switch v.Type {
				case jsonrpc.Ticker:
					// log.Infof("ticker: %+v", v.Ticker)
					dataStruct.Ticker.Price = v.Ticker.LTP

				case jsonrpc.Executions:
					// log.Infof("exec: %+v", v.Executions)

				case jsonrpc.Orderbook:
          // log.Infof("board: %+v", v.Orderbook)

        /*
          # Private channels
        */
				case jsonrpc.ChildOrders:
          // log.Infof("child orders: %+v", v.ChildOrders)

				case jsonrpc.ParentOrders:
          // log.Infof("parent orders: %+v", v.ParentOrders)
          
        case jsonrpc.Error:
          // do something()

        default:

        }
      }
    }
  )}

  if eg.Wait();err != nil {
    log.Error(err)
  }
}
```





# Author
[@_numbP](https://twitter.com/_numbp)