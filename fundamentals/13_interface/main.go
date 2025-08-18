package main

import "fmt"

type Alias = struct {
	name string
	age  int
}
type AliasPtr = *struct {
	name string
	age  int
}

type Writer interface {
	write([]byte)
	Write([]byte)
}

type W struct{}

func (w W) Write() {
	fmt.Println("write func call")
	// do something here
}

func main() {

	w := W{}
	w.Write()

	var a interface{} = (*int)(nil)
	var b interface{} = nil

	_ = b

	fmt.Println(a == nil)

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
