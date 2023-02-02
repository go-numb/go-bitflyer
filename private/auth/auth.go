package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Client struct {
	key    string
	secret string
}

func New(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
	}
}

func (p *Client) Get() (key, secret string) {
	return p.key, p.secret
}

func (p *Client) Signature(body string) string {
	mac := hmac.New(sha256.New, []byte(p.secret))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func (p *Client) WsParamForPrivate() (now int, nonce, sign string) {
	mac := hmac.New(sha256.New, []byte(p.secret))

	t := time.Now()
	now = int(t.Unix())
	nonce = fmt.Sprintf("%d", t.UnixNano())

	mac.Write([]byte(fmt.Sprintf("%d", now)))
	mac.Write([]byte(nonce))

	sign = hex.EncodeToString(mac.Sum(nil))
	return now, nonce, sign
}
