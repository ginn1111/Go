package main

import (
	"fmt"
)

func trySend(ch chan int, value int) bool {
	select {
	case ch <- value:
		return true
	default:
		fmt.Println("SEND FAILED!")
		return false
	}
}

func test() {
	ch := make(chan int)
	ch <- 42
	close(ch)
}

func main() {

	test()

}
