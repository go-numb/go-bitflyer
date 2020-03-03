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
- Performance tuned

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
  done := make(chen struct{})

  recieve := make(chan WsWriter)
	ctx, cancel := context.WithCancel(context.Background())

  go jsonrpc.Connect(ctx, recieve, []string{
		"lightning_board_snapshot",
		"lightning_board",
		"lightning_ticker",
		"lightning_executions",
	}, []string{
		string(types.FXBTCJPY),
		string(types.BTCJPY),
		string(types.ETHJPY),
  })
  
  go jsonrpc.ConnectForPrivate(ctx, recieve, <API_KEY>, <API_SECRET>, []string{
		"child_order_events",
		"parent_order_events",
	})

	go func() {
		for {
			select {
			case v := <-recieve:
				switch v.Types {
				case Board:
					fmt.Printf("%s - %+v\n", v.ProductCode, v.Board)
				case Ticker:
					fmt.Printf("%s - %+v\n", v.ProductCode, v.Ticker)
				case Executions:
					fmt.Printf("%s - %+v\n", v.ProductCode, v.Executions)

				case ChildOrders:
					fmt.Printf("%+v\n", v.ChildOrderEvent)
				case ParentOrders:
					fmt.Printf("%+v\n", v.ParentOrderEvent)

				case Undefined:
					fmt.Printf("undefined: %s - %+v\n", v.ProductCode, v.Results)
				case Error:
					fmt.Printf("error: %s - %+v\n", v.ProductCode, v.Results)

				}
			}

			select {
			case <-ctx.Done():
				return
			default:
			}
		}()

  <-done
}
```





# Author
[@_numbP](https://twitter.com/_numbp)