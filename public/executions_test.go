package public

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestToTime(t *testing.T) {
	j := `{
        "id": 2433730436,
        "side": "BUY",
        "price": 180000,
        "size": 0.1,
        "exec_date": "2017-02-05T17:58:33.59",
        "buy_child_order_acceptance_id": "JRF20170205-175833-006184",
        "sell_child_order_acceptance_id": "JRF20170205-175832-028676"
    }`

	var e Execution
	json.Unmarshal([]byte(j), &e)
	fmt.Println(e)
	times := e.ToDate()
	if times == nil {
		t.Fatal("results nil")
	}
	fmt.Printf("%s, %s\n", times[0], times[1])
}
