package private

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Withdraw struct {
	// CurrencyCode 現在は "JPY" のみ
	CurrencyCode  string  `json:"currency_code"`
	BankAccountID int     `json:"bank_account_id"`
	Amount        float64 `json:"amount"`
	// Code 二段階認証の確認コード
	Code string `json:"code,omitempty"`
}

type ResponseForWithdraw struct {
	MessageID string `json:"message_id,omitempty"`

	// when error
	Status       int         `json:"status,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (req *Withdraw) IsPrivate() bool {
	return true
}

func (req *Withdraw) Path() string {
	return "me/withdraw"
}

func (req *Withdraw) Method() string {
	return http.MethodPost
}

func (req *Withdraw) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Withdraw) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	return b
}
