package main

import "fmt"

type T int

func (t *T) M() {
	*t = 10
}

func Modify() {
	var t T = 1
	var p = &t
	var tP T = 2

	p.M()
	tP.M()

	fmt.Println(t)
	fmt.Println(tP)
}
