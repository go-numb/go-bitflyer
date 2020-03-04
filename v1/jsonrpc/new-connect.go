package jsonrpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-numb/go-bitflyer/v1/public/board"
	"github.com/go-numb/go-bitflyer/v1/public/executions"
	"github.com/go-numb/go-bitflyer/v1/public/ticker"
	jsoniter "github.com/json-iterator/go"

	"github.com/go-numb/go-bitflyer/v1/types"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
)

const (
	USE1                       = "wss://ws.lightstream.bitflyer.com/json-rpc"
	READDEADLINE time.Duration = 300
)

type Types int

const (
	All Types = iota
	Ticker
	Executions
	Board
	ChildOrders
	ParentOrders
	Undefined
	Error
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Request struct {
	Jsonrpc string                 `json:"jsonrpc,omitempty"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      int                    `json:"id,omitempty"`
}

func requests(conn *websocket.Conn, channels, symbols []string) (requests []Request, err error) {
	if symbols != nil {
		for i := range channels {
			for j := range symbols {
				requests = append(requests, Request{
					Jsonrpc: "2.0",
					Method:  "subscribe",
					Params: map[string]interface{}{
						"channel": fmt.Sprintf("%s_%s", channels[i], symbols[j]),
					},
					ID: 1,
				})
			}
		}
	} else {
		for i := range channels {
			requests = append(requests, Request{
				Jsonrpc: "2.0",
				Method:  "subscribe",
				Params: map[string]interface{}{
					"channel": channels[i],
				},
				ID: 1,
			})
		}
	}

	fmt.Printf("%+v\n", requests)

	for i := range requests {
		if err := conn.WriteJSON(requests[i]); err != nil {
			return nil, err
		}
	}

	return requests, nil
}

func unsubscribe(conn *websocket.Conn, requests []Request) {
	for i := range requests {
		requests[i].Method = "unsubscribe"
		if err := conn.WriteJSON(requests[i]); err != nil {
			fmt.Printf("%+v\n", err)
		}
	}

	defer conn.Close()
	fmt.Println("kill subscribed")
}

type WsWriter struct {
	Types       Types
	ProductCode types.ProductCode
	Board       board.Response
	Ticker      ticker.Response
	Executions  []executions.Execution

	ChildOrderEvent  []WsChildOrderEvent
	ParentOrderEvent []WsParentEvent

	Results error
}

func Connect(ctx context.Context, ch chan WsWriter, channels, symbols []string, log *logrus.Logger) {
	if log == nil {
		log = logrus.New()
	}

RECONNECT:
	conn, _, err := websocket.DefaultDialer.Dial(USE1, nil)
	if err != nil {
		log.Fatal(err)
	}

	requests, err := requests(conn, channels, symbols)
	if err != nil {
		log.Fatalf("disconnect %v", err)
	}

	var eg errgroup.Group
	eg.Go(func() error {
		for {
			conn.SetReadDeadline(time.Now().Add(READDEADLINE * time.Second))
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("can't receive error: %v", err)
			}
			// start := time.Now()

			name, err := jsonparser.GetString(msg, "params", "channel")
			if err != nil {
				continue
			}

			data, _, _, err := jsonparser.Get(msg, "params", "message")
			if err != nil {
				continue
			}

			var w WsWriter

			switch {
			case strings.HasPrefix(name, "lightning_board_snapshot_"):
				w.Types = Board
				if err := json.Unmarshal(data, &w.Board); err != nil {
					continue
				}

			case strings.HasPrefix(name, "lightning_board_"):
				w.Types = Board
				if err := json.Unmarshal(data, &w.Board); err != nil {
					continue
				}

			case strings.HasPrefix(name, "lightning_ticker_"):
				w.Types = Ticker
				if err := json.Unmarshal(data, &w.Ticker); err != nil {
					continue
				}

			case strings.HasPrefix(name, "lightning_executions_"):
				w.Types = Executions
				if err := json.Unmarshal(data, &w.Executions); err != nil {
					continue
				}

			default:
				w.Types = Undefined
				w.Results = fmt.Errorf("%v", string(msg))
			}

			switch { // switch with ProductCode
			case strings.HasSuffix(name, string(types.FXBTCJPY)):
				w.ProductCode = types.FXBTCJPY

			case strings.HasSuffix(name, string(types.BTCJPY)):
				w.ProductCode = types.BTCJPY

			case strings.HasSuffix(name, string(types.ETHJPY)):
				w.ProductCode = types.ETHJPY

			case strings.HasSuffix(name, string(types.ETHJPY)):
				w.ProductCode = types.ETHBTC
			default:
				w.ProductCode = types.UNDEFINED
			}

			select { // 外部からの停止
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			// log.Debugf("recieve to send time: %v\n", time.Now().Sub(start))
			ch <- w
		}
	})

	if err := eg.Wait(); err != nil {
		log.Errorf("%v", err)

		// 外部からのキャンセル
		if strings.Contains(err.Error(), context.Canceled.Error()) {
			// defer close()/unsubscribe()
			return
		}
	}

	// 明示的 Unsubscribed
	// context.cancel()された場合は
	unsubscribe(conn, requests)

	// Maintenanceならば待機
	// Maintenanceでなければ、即再接続
	if isMentenance() {
		for {
			if !isMentenance() {
				break
			}
			time.Sleep(time.Second)
		}
	}

	goto RECONNECT
}

func isMentenance() bool {
	// ServerTimeを考慮し、UTC基準に
	hour := time.Now().UTC().Hour()
	if hour != 19 {
		return false
	}

	if 12 < time.Now().Minute() { // メンテナンス以外
		return false
	}
	return true
}
