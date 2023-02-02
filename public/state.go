package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Status struct {
	ProductCode string `url:"product_code,omitempty"`
}

type ResponseForStatus struct {
	Health string `json:"health"`
	State  string `json:"state"`
	Data   struct {
		SpecialQuotation int `json:"special_quotation,omitempty"`
	} `json:"data,omitempty"`
}

func (req *Status) IsPrivate() bool {
	return false
}

func (req *Status) Path() string {
	return "getboardstate"
}

func (req *Status) Method() string {
	return http.MethodGet
}

func (req *Status) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Status) Payload() []byte {
	return nil
}
