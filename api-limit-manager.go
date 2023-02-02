package bitflyer

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type APILimit struct {
	Reset time.Time
	// リセットまでの秒数
	Period int
	// 残回数
	Remain int
}

func (p *APILimit) Set(header http.Header) {
	s := header.Get("X-Ratelimit-Reset")
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Print("headers has not key, ", err)
		return
	}
	p.Reset = time.Unix(i64, 0)

	s = header.Get("X-Ratelimit-Period")
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Print("headers has not key, ", err)
		return
	}
	p.Period = i

	s = header.Get("X-Ratelimit-Remaining")
	i, err = strconv.Atoi(s)
	if err != nil {
		log.Print("headers has not key, ", err)
		return
	}
	p.Remain = i
}

func (p *APILimit) IsRemain() bool {
	return p.Remain > 1
}

func (p *APILimit) Sleep() {
	time.Sleep(time.Duration(p.Period) * time.Second)
}
