package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Withdrawals struct {
	Count  int   `url:"count,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
	// MessageID 戻り値の受付 ID を指定して、出金状況を確認
	MessageID string `url:"message_id,omitempty"`
}

type ResponseForWithdrawals []Withdrawal

type Withdrawal struct {
	ID           int     `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func (req *Withdrawals) IsPrivate() bool {
	return true
}

func (req *Withdrawals) Path() string {
	return "me/getwithdrawals"
}

func (req *Withdrawals) Method() string {
	return http.MethodGet
}

func (req *Withdrawals) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Withdrawals) Payload() []byte {
	return nil
}
