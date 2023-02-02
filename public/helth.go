package public

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Helth struct {
	ProductCode string `url:"product_code,omitempty"`
}

type ResponseForHelth struct {
	Status string `json:"status"`
}

func (req *Helth) IsPrivate() bool {
	return false
}

func (req *Helth) Path() string {
	return "gethealth"
}

func (req *Helth) Method() string {
	return http.MethodGet
}

func (req *Helth) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Helth) Payload() []byte {
	return nil
}
