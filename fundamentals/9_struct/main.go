package main

import "fmt"

type Book struct {
	Page    int
	Content string
}

type Book2 struct {
	Content string
	Page    int
}

func main() {
	book := Book{Page: 10, Content: "Hello"}
	book2 := Book2{Page: 10, Content: "Hello"}

var config = struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080
}
}
