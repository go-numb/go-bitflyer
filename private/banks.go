package private

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Banks struct {
}

type ResponseForBanks []Bank

type Bank struct {
	ID            int    `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountType   string `json:"account_type"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

func (req *Banks) IsPrivate() bool {
	return true
}

func (req *Banks) Path() string {
	return "me/getbankaccounts"
}

func (req *Banks) Method() string {
	return http.MethodGet
}

func (req *Banks) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Banks) Payload() []byte {
	return nil
}
