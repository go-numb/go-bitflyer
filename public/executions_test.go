package public

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	j = `{
        "id": 2433730436,
        "side": "BUY",
        "price": 180000,
        "size": 0.1,
        "exec_date": "2017-02-05T17:58:33.59",
        "buy_child_order_acceptance_id": "JRF20170205-175833-006184",
        "sell_child_order_acceptance_id": "JRF20170205-175832-028676"
    }`
)

func TestToTime(t *testing.T) {
	var e Execution
	json.Unmarshal([]byte(j), &e)
	fmt.Println(e)
	times := e.ToDate()
	if times == nil {
		t.Fatal("results nil")
	}
	fmt.Printf("%s, %s\n", times[0].String(), times[1].String())

	assert.Equal(t, []string{"2017-02-05 17:58:33 +0000 UTC", "2017-02-05 17:58:32 +0000 UTC"}, []string{times[0].String(), times[1].String()}, "not match the correct")
}

func TestToUniqueID(t *testing.T) {
	var e Execution
	json.Unmarshal([]byte(j), &e)

	assert.Equal(t, "JRF20170205-175833-006184_JRF20170205-175832-028676", e.ToUniqueID(), "not match the correct")
}
