package circuitbreaker

import (
	"context"
	"fmt"
	"time"

	"github.com/go-numb/go-bitflyer/private/auth"
	"github.com/go-numb/go-bitflyer/public"
	"github.com/go-numb/go-bitflyer/realtime"
	"github.com/go-numb/go-bitflyer/types"
)

type Client struct {
	GetNotif chan *Notif
}

type Notif struct {
	IsBreak         bool
	Tick            public.ResponseForTicker
	CreatedServerAt time.Time
}

// Executor 監視し、本体に通知する
func (p *Client) Executor(ctx context.Context, symbols ...string) (err error) {
	ws := realtime.New(ctx)
	reciver := make(chan realtime.Response, 3)
	go ws.Connect(&auth.Client{}, symbols, []string{
		realtime.Ticker,
	}, reciver)

Exit:

	for {
		select {
		case v := <-reciver:
			switch v.Types {
			case realtime.TickerN:
				if v.Ticker.State != types.CIRCUITBREAK {
					continue
				}

				p.GetNotif <- &Notif{
					IsBreak:         true,
					Tick:            v.Ticker,
					CreatedServerAt: time.Now(),
				}
			}

		case <-ctx.Done():
			err = fmt.Errorf("context done: %s", ctx.Err())
			break Exit
		}
	}

	return err
}
