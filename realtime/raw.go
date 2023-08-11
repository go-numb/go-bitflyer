package realtime

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/rluisr/go-bitflyer/private/auth"
)

// If you want to utilize raw data
func (p *Client) Copy() *websocket.Conn {
	return p.conn
}

func (p *Client) Subscribe(conf *auth.Client, symbols, channels []string) []*Request {
	requests := _createRequester(conf, channels, symbols)
	if err := p.subscribe(
		conf,
		requests,
	); err != nil {
		log.Fatal("[FATAL] in proccess of subscribe, ", err.Error())
	}

	return requests
}

func (p *Client) Unsubscribe(requests []*Request) error {
	return p.unsubscribe(requests)
}
