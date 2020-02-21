package jsonrpc

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/sync/errgroup"

	"github.com/buger/jsonparser"
	"github.com/go-numb/go-bitflyer/auth"
	"github.com/go-numb/go-bitflyer/v1/public/ticker"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WsRequestForJSONRPC struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  WsParam `json:"params"`
	ID      int     `json:"id"`
}

type WsParam struct {
	APIKey    string `json:"api_key"`
	Timestamp int    `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

type WsResponseForAuth struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  bool   `json:"result"`
}

type WSResponseForTicker struct {
	ProductCode     string    `json:"product_code"`
	Timestamp       time.Time `json:"timestamp"`
	TickID          int       `json:"tick_id"`
	BestBid         float64   `json:"best_bid"`
	BestAsk         float64   `json:"best_ask"`
	BestBidSize     float64   `json:"best_bid_size"`
	BestAskSize     float64   `json:"best_ask_size"`
	TotalBidDepth   float64   `json:"total_bid_depth"`
	TotalAskDepth   float64   `json:"total_ask_depth"`
	Ltp             float64   `json:"ltp"`
	Volume          float64   `json:"volume"`
	VolumeByProduct float64   `json:"volume_by_product"`
}

type WsResponseForChildEvent struct {
	ExecID                 int    `json:"exec_id"`
	ProductCode            string `json:"product_code"`
	ChildOrderID           string `json:"child_order_id"`
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
	ChildOrderType         string `json:"child_order_type"`

	EventDate  time.Time `json:"event_date"`
	EventType  string    `json:"event_type"`
	Side       string    `json:"side"`
	Price      int       `json:"price"`
	Size       float64   `json:"size"`
	ExpireDate string    `json:"expire_date"`

	// 新設分追記
	Reason     string  `json:"reason"`
	Commission float64 `json:"commission"`
	SFD        float64 `json:"sfd"`
}

type WsResponseForParentEvent struct {
	ProductCode             string    `json:"product_code"`
	ParentOrderID           string    `json:"parent_order_id"`
	ParentOrderAcceptanceID string    `json:"parent_order_acceptance_id"`
	EventDate               time.Time `json:"event_date"`
	EventType               string    `json:"event_type"`
	ParentOrderType         string    `json:"parent_order_type"`
	Reason                  string    `json:"reason"`
	ParameterIndex          int       `json:"parameter_index"`
	ChildOrderType          string    `json:"child_order_type"`
	Side                    string    `json:"side"`
	Price                   int       `json:"price"`
	Size                    float64   `json:"size"`
	ExpireDate              time.Time `json:"expire_date"`
	ChildOrderAcceptanceID  string    `json:"child_order_acceptance_id"`
}

// GetPrivate is connect websocket, private channels
func GetPrivate(key, secret string, channels []string, ch chan Response) {
	conn, _, err := websocket.DefaultDialer.Dial(BASEURL, nil)
	if err != nil {
		ch <- Response{
			Type:  Error,
			Error: err,
		}
		return
	}
	defer conn.Close()

	if err := subscriber(conn, key, secret); err != nil {
		ch <- Response{
			Type:  Error,
			Error: err,
		}
		return
	}

	if err := writer(conn, channels); err != nil {
		ch <- Response{
			Type:  Error,
			Error: err,
		}
		return
	}

	var eg errgroup.Group
	eg.Go(func() error {
		for {
			conn.SetReadDeadline(time.Now().Add(HeartbeatIntervalSecond * 10 * time.Second))
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return errors.Wrap(err, "can't receive error: ")
			}

			name, err := jsonparser.GetString(msg, "params", "channel")
			if err != nil {
				continue
			}
			data, _, _, err := jsonparser.Get(msg, "params", "message")
			if err != nil {
				continue
			}

			switch name {
			case "lightning_ticker_BTC_JPY":
				// fmt.Printf("%+v\n", string(data))
				// SetDeadLine回避捨てイベント
				var parent ticker.Response
				json.Unmarshal(data, &parent)
				ch <- Response{
					Type:   Ticker,
					Ticker: parent,
				}

			case "lightning_ticker_FX_BTC_JPY":
				// SetDeadLine回避捨てイベント
				var parent ticker.Response
				json.Unmarshal(data, &parent)
				ch <- Response{
					Type:   Ticker,
					Ticker: parent,
				}

			case "child_order_events":
				var child []WsResponseForChildEvent
				json.Unmarshal(data, &child)
				ch <- Response{
					Type:        ChildOrders,
					ChildOrders: child,
				}

			case "parent_order_events":
				var parent []WsResponseForParentEvent
				json.Unmarshal(data, &parent)
				ch <- Response{
					Type:         ParentOrders,
					ParentOrders: parent,
				}

			default:
				ch <- Response{
					Type:  Error,
					Error: errors.New("read type error at private channel: " + name),
				}
			}
		}

	})

	if err := eg.Wait(); err != nil {
		log.Error(err)
		go func() {
			ch <- Response{
				Type:  Error,
				Error: errors.New("websocket type error: " + err.Error()),
			}
		}()
	}
}

func subscriber(conn *websocket.Conn, key, secret string) error {
	now, nonce, sign := auth.WsParamForPrivate(secret)
	req := &WsRequestForJSONRPC{
		Jsonrpc: "2.0",
		Method:  "auth",
		Params: WsParam{
			APIKey:    key,
			Timestamp: now,
			Nonce:     nonce,
			Signature: sign,
		},
		ID: now,
	}

	if err := conn.WriteJSON(req); err != nil {
		return err
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	t, _ := jsonparser.GetBoolean(msg, "result")
	if !t { // read channel return, if result  false
		return err
	}
	fmt.Printf("private channel connect success: %+v\n", t)

	return nil
}

func writer(conn *websocket.Conn, channels []string) error {
	var requests []string
	for _, channel := range channels {
		fmt.Println(channel)
		switch {
		case strings.HasPrefix(channel, "lightning_ticker_BTC_JPY"):
			fmt.Println("type has lightning_ticker_BTC_JPY")
		case strings.HasPrefix(channel, "lightning_ticker_FX_BTC_JPY"):
			fmt.Println("type has lightning_ticker_FX_BTC_JPY")
		case strings.HasPrefix(channel, "child_order_events"):
			fmt.Println("type has child order")
		case strings.HasPrefix(channel, "parent_order_events"):
			fmt.Println("type has parent order")
		}
		requests = append(requests, fmt.Sprintf(`{"jsonrpc": "2.0", "method": "subscribe", "params": {"channel": "%s"}, "id": %d}`, channel, time.Now().UTC().Unix()))
	}

	for _, v := range requests {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(v)); err != nil {
			return err
		}
	}

	return nil
}
