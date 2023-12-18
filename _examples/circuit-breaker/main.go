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

	<-ctx.Done()
}
