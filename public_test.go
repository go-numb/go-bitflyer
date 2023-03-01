package bitflyer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go-numb/go-bitflyer/public"
	"github.com/stretchr/testify/assert"

	"github.com/tcnksm/go-httpstat"
)

func TestRequestClientResolver(t *testing.T) {

	var (
		urls = []string{
			V1,
			V1 + "getmarkets",
			V1 + "board",
		}
	)

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			// DisableKeepAlives: true,
		},
	}
	log.Println("start reduce")
	for _, url := range urls {
		log.Printf("GET %s", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		result := new(httpstat.Result)
		ctx := httpstat.WithHTTPStat(req.Context(), result)
		req = req.WithContext(ctx)

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		result.End(time.Now())
		log.Printf("%+v\n", result)
	}

	log.Println("start new client")
	for _, url := range urls {
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				// DisableKeepAlives: true,
			},
		}
		log.Printf("GET %s", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		result := new(httpstat.Result)
		ctx := httpstat.WithHTTPStat(req.Context(), result)
		req = req.WithContext(ctx)

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		result.End(time.Now())
		log.Printf("%+v\n", result)
	}

}

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
