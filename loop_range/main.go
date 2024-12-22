package main

import "fmt"

type Student = struct {
	age  int
	name string
}

func main() {
	arr := []Student{
		{1, "a"},
	}

	for _, val := range arr {
		fmt.Println(val)
		val.age = 18
	}

	fmt.Println(arr[0])

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	i := 0

	fmt.Println(sum)

	sum = 0

	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum)
}
