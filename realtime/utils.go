package realtime

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"github.com/rluisr/go-bitflyer/private/auth"
)

func (p *Client) subscribe(conf *auth.Client, requests []*Request) error {
	log.Print("start subscribe------------")
	if conf != nil {
		if err := p.conn.WriteJSON(_auth(conf)); err != nil {
			return err
		}

		_, msg, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return err
			}
			log.Printf("[Ignore ERROR] can't receive error: %v", err)
		}

		if err := success("auth", msg); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	for i := range requests {
		_, isThere := requests[i].Params[CHANNEL].(string)
		if !isThere {
			continue
		}

		if err := p.conn.WriteJSON(requests[i]); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	log.Println("--ed")

	return nil
}

func (p *Client) unsubscribe(requests []*Request) error {
	log.Print("start [un]subscribe------------")
	for i := range requests {
		requests[i].Method = "unsubscribe"
		if err := p.conn.WriteJSON(requests[i]); err != nil {
			log.Printf("unsubscribed error: %s", err.Error())
			return err
		}
	}

	log.Println("--ed")
	return nil
}

func _createRequester(conf *auth.Client, channels, symbols []string) []*Request {
	req := make([]*Request, 0)

	for i := 0; i < len(channels); i++ {
		for j := 0; j < len(symbols); j++ {
			req = append(req, &Request{
				Jsonrpc: JSONRPCV,
				Method:  SUBSCRIBE,
				Params: map[string]interface{}{
					CHANNEL: fmt.Sprintf("lightning_%s_%s", channels[i], symbols[j]),
				},
				ID: i,
			})
		}
	}

	if conf != nil {
		req = append(req, &Request{
			Jsonrpc: JSONRPCV,
			Method:  SUBSCRIBE,
			Params: map[string]interface{}{
				CHANNEL: ParentOrders,
			},
			ID: len(req),
		})

		req = append(req, &Request{
			Jsonrpc: JSONRPCV,
			Method:  SUBSCRIBE,
			Params: map[string]interface{}{
				CHANNEL: ChildOrders,
			},
			ID: len(req),
		})
	}

	log.Println("created request parameters")
	for i, v := range req {
		channelname, isThere := v.Params[CHANNEL].(string)
		if !isThere {
			channelname = "auth signature"
		}
		fmt.Printf("%d: %d - %s\n", i, v.ID, channelname)
	}

	return req
}

func _auth(conf *auth.Client) *Request {
	key, _ := conf.Get()
	now, nonce, sign := conf.WsParamForPrivate()
	return &Request{
		Jsonrpc: JSONRPCV,
		Method:  AUTH,
		Params: map[string]interface{}{
			"api_key":   key,
			"timestamp": now,
			"nonce":     nonce,
			"signature": sign,
		},
		ID: now,
	}
}

func success(channelname string, msg []byte) error {
	id, _ := jsonparser.GetInt(msg, "id")
	result, _ := jsonparser.GetBoolean(msg, "result")
	_, _, _, err := jsonparser.Get(msg, "params")
	if err == nil {
		result = true
	}
	errmsg, err := jsonparser.GetString(msg, "error")
	if err == nil && errmsg != "" {
		return fmt.Errorf("[ERROR] get error param, id: %d, channel: %s, err msg: %s", id, channelname, errmsg)
	}

	if !result {
		return fmt.Errorf("[ERROR] subscribed error, id: %d, channel: %s, return %v, %s", id, channelname, result, errmsg)
	}

	log.Printf("[INFO] subscribed,  id: %d, subscribed: %s, success: %t\n", id, channelname, result)

	return nil
}

func _checkWebsocketErr(err error, errStruct interface{}) error {
	if err := json.Unmarshal([]byte(err.Error()), errStruct); err != nil {
		return fmt.Errorf("[ERROR] unmarshal error, %s", err.Error())
	}
	return nil
}
