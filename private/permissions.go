package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Permissions struct {
}

type ResponseForPermissions []string

func (req *Permissions) IsPrivate() bool {
	return true
}

func (req *Permissions) Path() string {
	return "me/getpermissions"
}

func (req *Permissions) Method() string {
	return http.MethodGet
}

func (req *Permissions) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Permissions) Payload() []byte {
	return nil
}
