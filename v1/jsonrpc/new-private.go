package jsonrpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-numb/go-bitflyer/auth"
	"github.com/sirupsen/logrus"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
)

func requestsForPrivate(conn *websocket.Conn, key, secret string) error {
	now, nonce, sign := auth.WsParamForPrivate(secret)
	req := &Request{
		Jsonrpc: "2.0",
		Method:  "auth",
		Params: map[string]interface{}{
			"api_key":   key,
			"timestamp": now,
			"nonce":     nonce,
			"signature": sign,
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
	isSuccess, _ := jsonparser.GetBoolean(msg, "result")
	if !isSuccess { // read channel return, if result  false
		return err
	}
	fmt.Printf("private channel connect success: %t\n", isSuccess)

	return nil
}

type WsChildOrderEvent struct {
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

type WsParentEvent struct {
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

func ConnectForPrivate(ctx context.Context, ch chan WsWriter, key, secret string, channels []string, log *logrus.Logger) {
	if log == nil {
		log = logrus.New()
	}

RECONNECT:
	conn, _, err := websocket.DefaultDialer.Dial(USE1, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := requestsForPrivate(conn, key, secret); err != nil {
		log.Fatalf("cant connect to private %v", err)
	}

	requests, err := requests(conn, channels, nil)
	if err != nil {
		log.Fatalf("disconnect %v", err)
	}
	defer unsubscribe(conn, requests)

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
			case strings.HasPrefix(name, "lightning_ticker"):
				w.Types = Ticker
				if err := json.Unmarshal(data, &w.Ticker); err != nil {
					continue
				}

			case strings.HasPrefix(name, "child_order_events"):
				w.Types = ChildOrders
				if err := json.Unmarshal(data, &w.ChildOrderEvent); err != nil {
					continue
				}

			case strings.HasPrefix(name, "parent_order_events"):
				w.Types = ParentOrders
				if err := json.Unmarshal(data, &w.ParentOrderEvent); err != nil {
					continue
				}

			default:
				w.Types = Undefined
				w.Results = fmt.Errorf("%v", string(msg))
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
		log.Errorf("bitflyer is mentenance time, in %s", time.Now().Format("2006/01/02 15:04:05"))
		for {
			if !isMentenance() {
				break
			}
			time.Sleep(time.Second)
		}
		log.Errorf("bitflyer mentenance is done, in %s", time.Now().Format("2006/01/02 15:04:05"))
	}

	goto RECONNECT
}
