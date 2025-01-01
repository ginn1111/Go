package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	go func(ch chan<- int, num int) {
		fmt.Println("Sending...")
		ch <- num
		fmt.Println("Sent")
	}(c, 3)

	done := make(chan struct{})

	go func(ch <-chan int) {
		time.Sleep(time.Second)
		fmt.Println("Receiving...")
		v := <-ch

		fmt.Println("Received:", v)
		done <- struct{}{}
	}(c)

	<-done

}
