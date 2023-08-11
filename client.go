package bitflyer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/rluisr/go-bitflyer/private/auth"
)

const (
	V1 = "https://api.bitflyer.com/v1/"
)

type Client struct {
	httpClient *http.Client
	// Auth
	Auth *auth.Client
}

func New(key, secret string, httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
		Auth:       auth.New(key, secret),
	}
}

type Require interface {
	IsPrivate() bool
	Path() string
	Method() string
	Query() string
	Payload() []byte
}

type Response struct {
	Code   int
	Status string
	Result interface{}

	Limit APILimit
	Error error
}

func (p *Client) request(req Require, results interface{}) (*APILimit, error) {
	res, err := p._do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// // Header
	// X-Frame-Options [SAMEORIGIN]
	// Strict-Transport-Security [max-age=31536000]
	// X-Ratelimit-Reset [1674714614]
	// X-Xss-Protection [1; mode=block]
	// Content-Security-Policy [default-src http: https: ws: wss: data: 'unsafe-inline' 'unsafe-eval'; img-src https: data: blob: 'self']
	// Vary [Accept-Encoding]
	// Date [Thu, 26 Jan 2023 06:25:14 GMT]
	// X-Ratelimit-Period [300]
	// X-Ratelimit-Remaining [499]
	// Content-Type [application/json; charset=utf-8]
	// Expires
	manage := new(APILimit)
	manage.Set(res.Header)

	// response change to result structure
	if err := _decode(res, results); err != nil {
		return nil, err
	}

	return manage, nil
}

func (p *Client) _do(req Require) (*http.Response, error) {
	r := p._newRequest(req)

	res, err := p.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	// no usefull headers
	if res.StatusCode != 200 {
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if strings.Contains(string(b), "Order not found") {
			return nil, fmt.Errorf("status: %s, order could not be found. check the designated ID", res.Status)
		}

		return nil, fmt.Errorf("status: %s, body: %q", res.Status, b)
	}

	return res, nil
}

func (p *Client) _newRequest(req Require) *http.Request {
	u, _ := url.Parse(V1)
	u.Path = path.Join(u.Path, req.Path())
	u.RawQuery = req.Query()

	r, err := http.NewRequest(req.Method(), u.String(), bytes.NewBuffer(req.Payload()))
	if err != nil {
		return nil
	}

	if req.IsPrivate() {
		nonce := time.Now().String()
		var path = u.Path
		if u.RawQuery != "" {
			path += "?" + u.RawQuery
		}
		payload := nonce + req.Method() + path + string(req.Payload())

		key, _ := p.Auth.Get()
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("ACCESS-KEY", key)
		r.Header.Set("ACCESS-SIGN", p._signature(payload))
		r.Header.Set("ACCESS-TIMESTAMP", nonce)
	}

	return r
}

func _decode(res *http.Response, out interface{}) error {
	// b, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("----\n%s\n----------\n", string(b))
	// fmt.Println(res.StatusCode)

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	return nil
}

func (p *Client) _signature(payload string) string {
	return p.Auth.Signature(payload)
}
