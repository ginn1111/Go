package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan *[16]byte, 1)

	c <- new([16]byte)
	close(c)

	// go func() {
	// 	dA := new([16]byte)
	//
	// 	for false {
	// 		_, err := rand.Read(dA[:])
	//
	// 		if err != nil {
	// 			close(c)
	// 		}
	//
	// 		// c <- dA
	// 		// dA, dB = dB, dA
	// 	}
	// }()

	for data := range c {
		fmt.Println(data)
	}

}
