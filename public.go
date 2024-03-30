package bitflyer

import (
	"github.com/go-numb/go-bitflyer/public"
)

func (p *Client) Status(req *public.Status) (*public.ResponseForStatus, *APILimit, error) {
	res := new(public.ResponseForStatus)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) Helth(req *public.Helth) (*public.ResponseForHelth, *APILimit, error) {
	res := new(public.ResponseForHelth)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return res, apiLimit, err
}

func (p *Client) Fr(req *public.Fr) (*public.ResponseForFr, *APILimit, error) {
	res := new(public.ResponseForFr)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return res, apiLimit, err
}

func (p *Client) LeverageC(req *public.LeverageC) (*public.ResponseForLeverageC, *APILimit, error) {
	res := new(public.ResponseForLeverageC)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return res, apiLimit, err
}

func (p *Client) Ticker(req *public.Ticker) (*public.ResponseForTicker, *APILimit, error) {
	res := new(public.ResponseForTicker)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return res, apiLimit, err
}

func (p *Client) Markets(req *public.Markets) ([]public.Market, *APILimit, error) {
	res := new(public.ResponseForMarkets)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return []public.Market(*res), apiLimit, err
}

func (p *Client) Board(req *public.Board) (*public.ResponseForBoard, *APILimit, error) {
	res := new(public.ResponseForBoard)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return res, apiLimit, err
}

func (p *Client) Executions(req *public.Executions) ([]public.Execution, *APILimit, error) {
	res := new(public.ResponseForExecutions)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return []public.Execution(*res), apiLimit, err
}

func (p *Client) Chat(req *public.Chat) ([]public.Comment, *APILimit, error) {
	res := new(public.ResponseForChat)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}
	return []public.Comment(*res), apiLimit, err
}
