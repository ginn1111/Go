package main

import (
	"fmt"
	"sync"
)

// distribute computation
func generator(in ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, i := range in {
			out <- i
		}
	}()

	return out
}

func cube(out <-chan int) <-chan int {
	tranOut := make(chan int)

	go func() {
		defer close(tranOut)

		for o := range out {
			tranOut <- o * o * o
		}
	}()

	return tranOut
}

func fanOut(in <-chan int, numOfWorker int) [](chan int) {
	workers := make([]chan int, numOfWorker)

	for i := range numOfWorker {
		out := make(chan int)
		workers[i] = out

		go func(ch chan<- int) {
			defer close(ch)

			for v := range in {
				fmt.Printf("Worker %d received: %d\n", i, v)
				ch <- v

			}
		}(workers[i])
	}

	return workers
}

func fanOut2(in <-chan int, numOfWorker int) []chan int {
	workers := make([]chan int, numOfWorker)

	for i := range numOfWorker {
		workers[i] = make(chan int)
	}

	go func() {
		i := 0
		for v := range in {
			workers[i] <- v
			fmt.Printf("Worker %d received: %d\n", i, v)
			i++
		}
		for _, worker := range workers {
			close(worker)
		}
	}()

	return workers
}

func fanIn(in ...(chan int)) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(in))

	for _, ci := range in {
		go func(ci <-chan int) {
			defer wg.Done()

			for v := range ci {
				out <- v
			}
		}(ci)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	return out
}

func main() {
	c1 := generator(1, 2, 3)

	trans1 := cube(c1)

	workers := fanOut2(trans1, 3)

	out := fanIn(workers...)

	sum := 0

	for v := range out {
		sum += v
	}

	fmt.Print(sum)
}
