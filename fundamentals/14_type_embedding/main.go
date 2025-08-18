package main

import "fmt"

type Book struct {
	Pages int
}

type PointerInt *int

type PInt PointerInt

type Int int

type PointerStruct struct {
	Int
}

func (b *Book) GetPages() int {
	return b.Pages
}

func main() {
	// literal struct type does not define a method
	// but it can own the method via embedding type
	a := struct {
		*Book
		name string
	}{
		&Book{123},
		"Thuan",
	}

	pA := &a

	fmt.Println(a.Book.GetPages(), pA.GetPages())

}
