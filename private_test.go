package bitflyer

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-numb/go-bitflyer/private"
)

const (
	KEY    = ""
	SECRET = ""
)

func init() {

}

func TestPermissions(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Permissions(&private.Permissions{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestBalance(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Balance(&private.Balance{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCollateral(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Collateral(&private.Collateral{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCollateralAccounts(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.CollateralAccounts(&private.CollateralAccounts{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestAddresses(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Addresses(&private.Addresses{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCoinins(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Coinins(&private.Coinins{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCoinouts(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Coinouts(&private.Coinouts{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestBanks(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Banks(&private.Banks{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestDeposits(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Deposits(&private.Deposits{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestWithdrawals(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Withdrawals(&private.Withdrawals{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestChildOrder(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.ChildOrder(&private.ChildOrder{
		ProductCode:    "FX_BTC_JPY",
		ChildOrderType: "LIMIT",
		Side:           "BUY",
		Price:          5_000_000,
		Size:           0.1,
		MinuteToExpire: 1000,
		TimeInForce:    "GTC",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCancelChildOrder(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.CancelChildOrder(&private.CancelChildOrder{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestParentOrder(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.ParentOrder(&private.ParentOrder{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCancelParentOrder(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.CancelParentOrder(&private.CancelParentOrder{
		ProductCode: "FX_BTC_JPY",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCancel(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Cancel(&private.Cancel{
		ProductCode: "FX_BTC_JPY",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestChildOrders(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.ChildOrders(&private.ChildOrders{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestDetailParentOrder(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.DetailParentOrder(&private.DetailParentOrder{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestMyExecutions(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.MyExecutions(&private.Executions{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestBalanceHistory(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.BalanceHistory(&private.BalanceHistory{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestPositions(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Positions(&private.Positions{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCollateralHistory(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.CollateralHistory(&private.CollateralHistory{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	fmt.Printf("%s\n", limit.Reset.String())
}

func TestCommission(t *testing.T) {
	key, secret := KEY, SECRET
	if key == "" {
		key = os.Getenv("bf_key")
	}
	if secret == "" {
		secret = os.Getenv("bf_secret")
	}

	client := New(key, secret)

	res, limit, err := client.Commission(&private.Commission{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

	fmt.Printf("%s\n", limit.Reset.String())
}
