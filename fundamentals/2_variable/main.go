package main

import "fmt"

func main() {
	var intTyped int = 1
	intInterTyped := 1
	var float32Typed float32 = 3.1
	var float64Typed float64 = 1 << 64

	converFloatFromInt := int(float32Typed)

	fmt.Print(intTyped, intInterTyped, float32Typed, float64Typed, converFloatFromInt)

}
