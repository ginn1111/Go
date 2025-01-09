package main

import "fmt"

type Language struct {
	Name    string
	Version float32
}

func main() {
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	copy(s[:4], s[:3])
	fmt.Println(s)

}
