package main

import "fmt"

type Student struct {
	id        string
	firstname string
	lastname  string
	dob       string
}

func (s Student) Fullname() string {
	return s.firstname + " " + s.lastname
}

func main() {
	s1 := Student{
		"1", "Pham Van", "Thuan", "13/09/2001",
	}

	pS1 := &Student{
		"2", "Nguyen Xuan", "Hoa", "24/01/2001",
	}

	fmt.Println(s1.Fullname(), (*pS1).Fullname())

}
