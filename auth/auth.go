package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-numb/go-bitflyer"
	"github.com/pkg/errors"
)

type AuthConfig struct {
	APIKey    string
	APISecret string
}

// SetAuthHeaders Signture to header
func SetAuthHeaders(config *AuthConfig, api api.API, req api.Request) (*http.Header, error) {
	url, err := api.ToURL()
	if err != nil {
		return nil, errors.Wrapf(err, "set base URL")
	}
	url.RawQuery = req.Query()

	method := req.Method()
	var path = url.Path
	if url.RawQuery != "" {
		path = url.Path + "?" + url.RawQuery
	}
	payload := req.Payload()

	mac := hmac.New(sha256.New, []byte(config.APISecret))
	// .jp -> Now()
	// .com -> UTC()
	t := time.Now().UTC().String()
	mac.Write([]byte(t))
	mac.Write([]byte(method))
	mac.Write([]byte(path))
	if len(payload) != 0 {
		mac.Write(payload)
	}

	sign := hex.EncodeToString(mac.Sum(nil))

	header := http.Header{}
	header.Set("ACCESS-KEY", config.APIKey)
	header.Set("ACCESS-TIMESTAMP", t)
	header.Set("ACCESS-SIGN", sign)

	return &header, nil
}

// WsParamForPrivate return util for private websocket
func WsParamForPrivate(sercret string) (now int, nonce, sign string) {
	mac := hmac.New(sha256.New, []byte(sercret))

	t := time.Now().UTC()
	rand.Seed(t.UnixNano())

	now = int(t.Unix())
	nonce = fmt.Sprintf("%d", rand.Int())

	mac.Write([]byte(fmt.Sprintf("%d%s", now, nonce)))

	sign = hex.EncodeToString(mac.Sum(nil))
	return now, nonce, sign
}
