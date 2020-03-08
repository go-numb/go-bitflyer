package v1

import (
	"net/http"
	"time"

	"github.com/go-numb/go-bitflyer/auth"
	"github.com/go-numb/go-bitflyer/httpclient"
	"github.com/go-numb/go-bitflyer/v1/private/addresses"
	"github.com/go-numb/go-bitflyer/v1/private/amounts"
	"github.com/go-numb/go-bitflyer/v1/private/balance"
	"github.com/go-numb/go-bitflyer/v1/private/bankaccounts"
	corders "github.com/go-numb/go-bitflyer/v1/private/cancels/orders"
	cpositions "github.com/go-numb/go-bitflyer/v1/private/cancels/positions"
	"github.com/go-numb/go-bitflyer/v1/private/childorders"
	"github.com/go-numb/go-bitflyer/v1/private/coinouts"
	"github.com/go-numb/go-bitflyer/v1/private/collateral"
	ex "github.com/go-numb/go-bitflyer/v1/private/executions"
	"github.com/go-numb/go-bitflyer/v1/private/histories"
	"github.com/go-numb/go-bitflyer/v1/private/orders/single"
	"github.com/go-numb/go-bitflyer/v1/private/orders/sp"
	"github.com/go-numb/go-bitflyer/v1/private/permissions"
	"github.com/go-numb/go-bitflyer/v1/private/positions"
	"github.com/go-numb/go-bitflyer/v1/public/board"
	"github.com/go-numb/go-bitflyer/v1/public/chats"
	"github.com/go-numb/go-bitflyer/v1/public/coinins"
	"github.com/go-numb/go-bitflyer/v1/public/executions"
	"github.com/go-numb/go-bitflyer/v1/public/health"
	"github.com/go-numb/go-bitflyer/v1/public/markets"
	"github.com/go-numb/go-bitflyer/v1/public/ticker"
	"github.com/pkg/errors"
)

const (
	APIHost    string = "https://api.bitflyer.jp"
	APIHostCom string = "https://api.bitflyer.com"
)

type Client struct {
	Host string

	HTTPClient *httpclient.Client

	AuthConfig *auth.AuthConfig
	// httpClient *http.Client
}

type ClientOpts struct {
	AuthConfig *auth.AuthConfig
}

func NewClient(opts *ClientOpts) *Client {
	return &Client{
		// Host:       APIHost,
		Host: APIHostCom,
		HTTPClient: httpclient.New(&http.Client{
			Timeout: 10 * time.Second,
		}, opts.AuthConfig),
		AuthConfig: opts.AuthConfig,
	}
}

func (c *Client) APIHost() string {
	return c.Host
}

/*
	# Public APIs
*/
// Markets
func (c *Client) Markets(req *markets.Request) (*markets.Response, *http.Response, error) {
	res := new(markets.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, markets.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Board(req *board.Request) (*board.Response, *http.Response, error) {
	res := new(board.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, board.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}

	return res, raw, nil
}

func (c *Client) Ticker(req *ticker.Request) (*ticker.Response, *http.Response, error) {
	res := new(ticker.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, ticker.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Executions(req *executions.Request) (*executions.Response, *http.Response, error) {
	res := new(executions.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, executions.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}

	return res, raw, nil
}

func (c *Client) Health(req *health.Request) (*health.Response, *http.Response, error) {
	res := new(health.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, health.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Chats(req *chats.Request) (*chats.Response, *http.Response, error) {
	res := new(chats.Response)
	raw, err := c.HTTPClient.Request(NewAPI(c, chats.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

/*
	# Private APIs
*/
// Permissions
func (c *Client) Permissions(req *permissions.Request) (*permissions.Response, *http.Response, error) {
	res := new(permissions.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, permissions.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Balance(req *balance.Request) (*balance.Response, *http.Response, error) {
	res := new(balance.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, balance.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Collateral(req *collateral.Request) (*collateral.Response, *http.Response, error) {
	res := new(collateral.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, collateral.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}

	return res, raw, nil
}

func (c *Client) CollateralAccounts(req *amounts.Request) (*amounts.Response, *http.Response, error) {
	res := new(amounts.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, amounts.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Addresses(req *addresses.Request) (*addresses.Response, *http.Response, error) {
	res := new(addresses.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, addresses.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Coinins(req *coinins.Request) (*coinins.Response, *http.Response, error) {
	res := new(coinins.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, coinins.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Coinouts(req *coinouts.Request) (*coinouts.Response, *http.Response, error) {
	res := new(coinouts.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, coinouts.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) BankAccounts(req *bankaccounts.Request) (*bankaccounts.Response, *http.Response, error) {
	res := new(bankaccounts.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, bankaccounts.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) OrderSingle(req *single.Request) (*single.Response, *http.Response, error) {
	res := new(single.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, single.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) OrderSP(req *sp.Request) (*sp.Response, *http.Response, error) {
	res := new(sp.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, sp.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) CancelOrderAll(req *corders.Request) (*corders.Response, *http.Response, error) {
	res := new(corders.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, corders.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) CancelByID(req *cpositions.Request) (*cpositions.Response, *http.Response, error) {
	res := new(cpositions.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, cpositions.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Positions(req *positions.Request) (*positions.Response, *http.Response, error) {
	res := new(positions.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, positions.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) ExecutionsMe(req *ex.Request) (*ex.Response, *http.Response, error) {
	res := new(ex.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, ex.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) ChildOrdersMe(req *childorders.Request) (*childorders.Response, *http.Response, error) {
	res := new(childorders.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, childorders.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}

func (c *Client) Histories(req *histories.Request) (*histories.Response, *http.Response, error) {
	res := new(histories.Response)
	raw, err := c.HTTPClient.Auth().Request(NewAPI(c, histories.APIPath), req, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "sends request")
	}
	return res, raw, nil
}
