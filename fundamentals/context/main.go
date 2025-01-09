package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	cancel()

	// ctx, cancel := context.WithCancel(ctx)
	//
	// go func() {
	// 	sc := bufio.NewScanner(os.Stdin)
	// 	sc.Scan()
	// 	cancel()
	// }()

	sleepAndTalk(ctx, 5*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(d):
		fmt.Println(msg)
	}
}
