package main

import (
	"fmt"
	"time"
)

func main() {

	myZodiacSign := "virgo"

	if myZodiacSign == "virgo" {
		fmt.Println("Virgo")
	}

	var i int

	for i = range 10 {
		fmt.Println(i)
	}

	c := make(chan int)

	go func() {
		c <- 2
	}()

	ii := 0

	for ii > 10 {
		fmt.Println(ii)

		ii++
	}

	go func() {
		switch cc := <-c; cc % 2 {

		case 1:
			fmt.Print("Odd")

		case 0:
			fmt.Print("Even")

		default:
			fmt.Print("Default")
		}
	}()

	time.Sleep(1 * time.Second)
}
