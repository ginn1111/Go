package main

import rt "go/reflections/reflect_type"
import rv "go/reflections/reflect_value"
import "fmt"

func main() {
	rt.WithInterface()
	rt.WithNonInterface()

	fmt.Println()

	rv.WithInterface()
	rv.WithNoneInterface()
}
