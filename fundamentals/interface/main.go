package main

import "fmt"

func main() {

	var x = []any{
		struct {
			name string
			age  int
		}{name: "Thuan", age: 24},
		1, 2,
		"Go",
	}

	for _, v := range x {
		switch i := v.(type) {
		case string:
			fmt.Println("String type", i)
		case struct {
			name string
			age  int
		}:

			fmt.Println("Struct type", i)
		default:
			fmt.Println("Default", i)

		}
	}
}
