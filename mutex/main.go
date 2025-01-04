package main

import (
	"log"
	"sync"
	"time"
)

type IntType struct {
	val int
	sync.Mutex
}

// a simple function that returns true if a number is even
func (i IntType) isEven() bool {
	return i.val%2 == 0
}

func main() {
	n := IntType{val: 0}

	go func() {

		func() {

			n.Lock()
			defer n.Unlock()
			isEven := n.isEven()

			log.Println("Start Sleep")
			time.Sleep(time.Millisecond * 500)
			if isEven {
				log.Println(n.val, "is even")
			}
		}()

		time.Sleep(2 * time.Second)
		log.Println("AFTER Sleep 5s")
	}()

	go func() {
		n.Lock()
		defer n.Unlock()

		log.Println("HEEE")

		n.val++
	}()

}
