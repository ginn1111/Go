package main

import (
	"fmt"
	_ "fmt"
)

type Language struct {
	Name    string
	Version float32
}

func getSlice() []int {
	return []int{1, 2, 3}
}

func getArray() [4]int {
	return [...]int{1, 2, 3, 4}
}

type S []int
type A [2]int
type P *A

func main() {

	fmt.Println(&((*[5]int)(nil))[:][0])

}
