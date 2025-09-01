package main

import (
	"context"
	"fmt"
	"time"
)

const KEY = "KEY"

func f(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("DONE")

	case <-time.After(time.Second * 2):
		fmt.Println("TIME OUT", ctx.Value(KEY))

	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	ctx = context.WithValue(ctx, KEY, 123)

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	f(ctx)

}
