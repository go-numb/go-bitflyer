package main

import (
	"context"
	"log"
)

func main() {
	client := New()
	ctx := context.Background()
	go func() {
		if err := client.Executor(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

Exit:
	for {
		select {
		case v := <-client.GetNotif:
			if !v.IsBreak {
				continue
			}
			log.Printf("get notifiaction: %v", v)
			dosometing()

		case <-ctx.Done():
			break Exit
		}
	}

	log.Fatal("done")
}

func dosomething() {
	// Do!
}
