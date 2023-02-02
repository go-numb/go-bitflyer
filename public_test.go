package bitflyer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/go-numb/go-bitflyer/public"
	"github.com/stretchr/testify/assert"
)

func TestMarkets(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Markets(&public.Markets{})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestBoard(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Board(&public.Board{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("results: ", res.MidPrice)
	for i := 0; i < len(res.Asks); i++ {
		fmt.Println(res.Asks[i])
	}
	for i := 0; i < len(res.Bids); i++ {
		fmt.Println(res.Bids[i])
	}
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestExecutions(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Executions(&public.Executions{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("results: ", res)
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestStatus(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Status(&public.Status{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("results: ", res)
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestTicker(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Ticker(&public.Ticker{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("results: ", res)
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestChat(t *testing.T) {
	client := New("", "")

	res, limit, err := client.Chat(&public.Chat{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("results: ", res)
	fmt.Printf("%s\n", limit.Reset.String())
}

func TestPowerWorks(t *testing.T) {
	client := New("", "")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(client *Client) {
		defer wg.Done()

		res, limit, err := client.Executions(&public.Executions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res {
			fmt.Println(v)
		}
		fmt.Println(limit.Remain)
	}(client)

	wg.Add(1)
	go func(clients *Client) {
		defer wg.Done()

		res, limit, err := clients.Executions(&public.Executions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res {
			fmt.Println(v)
		}
		fmt.Println(limit.Remain)
	}(client)

	wg.Wait()

	fmt.Println("done")
}

func TestDoubleWorks(t *testing.T) {
	forManagePublic := New("", "")
	forManagePrivate := New("", "")

	clients := make(map[int]*Client)
	clients[0] = forManagePublic
	clients[1] = forManagePrivate

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(client *Client) {
		defer wg.Done()

		res, limit, err := client.Executions(&public.Executions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res {
			fmt.Println(v)
		}
		fmt.Println(limit.Remain)
	}(clients[0])

	wg.Add(1)
	go func(clients *Client) {
		defer wg.Done()

		res, limit, err := clients.Executions(&public.Executions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res {
			fmt.Println(v)
		}
		fmt.Println(limit.Remain)
	}(clients[1])

	wg.Wait()

	fmt.Println("done")
}

func TestUnmarshal(t *testing.T) {
	f, _ := os.Open("./_data/response.json")
	defer f.Close()

	var r = new(public.ResponseForBoard)

	if err := json.NewDecoder(f).Decode(r); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, r.MidPrice, 2999406.0)

	t.Log("success")
}
