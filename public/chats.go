package public

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Chat struct {
	FromDate time.Time `url:"from_date,omitempty"`
}

type ResponseForChat []Comment

type Comment struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

func (req *Chat) IsPrivate() bool {
	return false
}

func (req *Chat) Path() string {
	return "getchats"
}

func (req *Chat) Method() string {
	return http.MethodGet
}

func (req *Chat) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Chat) Payload() []byte {
	return nil
}
