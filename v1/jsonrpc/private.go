package jsonrpc

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-numb/go-bitflyer/auth"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WsResponceForAuth struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  bool   `json:"result"`
}

type WsResponceForChildEvent struct {
	ProductCode            string    `json:"product_code"`
	ChildOrderID           string    `json:"child_order_id"`
	ChildOrderAcceptanceID string    `json:"child_order_acceptance_id"`
	EventDate              time.Time `json:"event_date"`
	EventType              string    `json:"event_type"`
	ChildOrderType         string    `json:"child_order_type"`
	Side                   string    `json:"side"`
	Price                  int       `json:"price"`
	Size                   float64   `json:"size"`
	ExpireDate             string    `json:"expire_date"`
}

type WsResponceForParentEvent struct {
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
		panic(err)
	}
	defer conn.Close()

	now, apiKey, nonce, sign := auth.WsParamForPrivate(key, secret)
	req := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "auth", "params": {"api_key": "%s", "timestamp": %d, "nonce": "%v", "signature": "%s"}, "id": %d}`, apiKey, now, nonce, sign, now)
	fmt.Printf("%+v\n", req)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(req)); err != nil {
		panic(err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	var check WsResponceForAuth
	json.Unmarshal(msg, &check)

	if !check.Result { // read channel return, if result  false
		panic(err)
	}
	fmt.Printf("private channel read success: %+v\n", check)

	var requests []string
	for _, channel := range channels {
		fmt.Println(channel)
		switch {
		case strings.HasPrefix(channel, "child_order_events"):
			fmt.Println("type has child order")
		case strings.HasPrefix(channel, "parent_order_events"):
			fmt.Println("type has parent order")
		}
		requests = append(requests, fmt.Sprintf(`{"jsonrpc": "2.0", "method": "subscribe", "params": {"channel": "%s"}, "id": %d}`, channel, time.Now().UTC().Unix()))
	}

	if len(requests) == 2 {
		fmt.Println("gets all private channels")
	}

	for _, v := range requests {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(v)); err != nil {
			panic(err)
		}
	}

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Minute))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			goto EXIT
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
		case "child_order_events":
			var child []WsResponceForChildEvent
			json.Unmarshal(data, &child)
			ch <- Response{
				Type:        ChildOrders,
				ChildOrders: child,
			}

		case "parent_order_events":
			var parent []WsResponceForParentEvent
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

EXIT:
}
