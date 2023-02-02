package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type CollateralAccounts struct {
}

type ResponseForCollateralAccounts []Currency

func (req *CollateralAccounts) IsPrivate() bool {
	return true
}

func (req *CollateralAccounts) Path() string {
	return "me/getcollateralaccounts"
}

func (req *CollateralAccounts) Method() string {
	return http.MethodGet
}

func (req *CollateralAccounts) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *CollateralAccounts) Payload() []byte {
	return nil
}
