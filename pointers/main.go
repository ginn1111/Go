package main

import (
	. "fmt"
)

type MyInt int
type Ta *int
type Tb *MyInt

func main() {
	var _ Ta
	var _ Ta
	var t MyInt = 1
	var t3 Ta
	var t4 *MyInt = &t

	Println(t3)
	Println(t3 == (*int)(t4), t3 == nil)
	Println(Tb(t4))

}
