package main

import "fmt"

func Test1(a, b int) (int, int) {
	// empty

	return 1, 2
}

func _() {
	fmt.Println("blank")
}

func init() {}

func Test2(a ...int) int {

	return 2
}

func main() {
	fmt.Printf("%T", Test2)
}
