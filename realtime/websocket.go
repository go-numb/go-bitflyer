package realtime

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-numb/go-bitflyer/private/auth"
	"github.com/gorilla/websocket"

	"github.com/buger/jsonparser"
)

const (
	ENDPOINT                   = "wss://ws.lightstream.bitflyer.com/json-rpc"
	READDEADLINE time.Duration = 300 * time.Second
)

type Client struct {
	conn *websocket.Conn
	ctx  context.Context
}

func New(ctx context.Context) *Client {
	conn, _, err := websocket.DefaultDialer.Dial(ENDPOINT, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		conn: conn,
		ctx:  ctx,
	}
}

func (p *Client) Close() error {
	if err := p.conn.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Client) Connect(conf *auth.Client, channels, symbols []string, send chan Response) error {
	defer log.Println("defer is end, completed websocket connect")
	defer p.Close()

	requests := _createRequester(conf, channels, symbols)
	if err := p.subscribe(
		conf,
		requests,
	); err != nil {
		log.Fatal("[FATAL] ", err.Error())
	}
	defer p.unsubscribe(requests)

L:
	for {
		select {
		case <-p.ctx.Done():
			log.Println("recived context cancel from parent, websocket closed")
			break L

		default:
			res := new(Response)
			_, msg, err := p.conn.ReadMessage()
			if err != nil {
				var er = new(websocket.CloseError)
				if err := _checkWebsocketErr(err, er); err != nil {
					return err
				}

				if er.Code == websocket.CloseAbnormalClosure {
					continue
				}

				p.conn.WriteMessage(websocket.BinaryMessage, []byte(`ping`))

				send <- res._set(err)
				continue
			}

			channelname, err := jsonparser.GetString(msg, "params", "channel")
			if err != nil {
				send <- res._set(err)
				continue
			}
			data, _, _, err := jsonparser.Get(msg, "params", "message")
			if err != nil {
				send <- res._set(err)
				continue
			}

			switch {
			case strings.HasPrefix(channelname, Ticker):
				// fmt.Println(Ticker)
				res.Types = TickerN
				res.ProductCode = strings.Replace(channelname, Ticker, "", 1)
				if err := json.Unmarshal(data, &res.Ticker); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal ticker %+v", err))
				}
			case strings.HasPrefix(channelname, Executions):
				// fmt.Println(Executions)
				res.Types = ExecutionsN
				res.ProductCode = strings.Replace(channelname, Executions, "", 1)
				if err := json.Unmarshal(data, &res.Executions); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal executions %+v", err))
				}
			case strings.HasPrefix(channelname, BoardSnap):
				// fmt.Println(BoardSnap)
				res.Types = BoardSnapN
				res.ProductCode = strings.Replace(channelname, BoardSnap, "", 1)
				if err := json.Unmarshal(data, &res.Board); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal board snap %+v", err))
				}
			case strings.HasPrefix(channelname, Board):
				// fmt.Println(Board)
				res.Types = BoardN
				res.ProductCode = strings.Replace(channelname, Board, "", 1)
				if err := json.Unmarshal(data, &res.Board); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal board update %+v", err))
				}
			case strings.HasPrefix(channelname, ChildOrders):
				// fmt.Println(ChildOrders)
				res.Types = ChildOrdersN
				if err := json.Unmarshal(data, &res.ChildOrders); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal childorder %+v", err))
				}
			case strings.HasPrefix(channelname, ParentOrders):
				// fmt.Println(ParentOrders)
				res.Types = ParentOrdersN
				if err := json.Unmarshal(data, &res.ParentOrders); err != nil {
					res.Types = ErrorN
					res._set(fmt.Errorf("[WARN]: cant unmarshal parentorder %+v", err))
				}
			case strings.HasPrefix(channelname, Error):
				// fmt.Println("error!")
				res.Types = ErrorN
				res._set(err)
			default:
				// fmt.Println("undefined", data)
				res.Types = UndefinedN
			}

			send <- *res
		}
	}

	return p.ctx.Err()
}
