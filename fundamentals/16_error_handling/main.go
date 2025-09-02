package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type error interface {
	Error() string
}

type Error1 struct {
	text string
}

type Error2 struct {
	text string
}

func (e Error1) Error() string {
	return "Error 1: " + e.text
}

func (e Error2) Error() string {
	return "Error 2: " + e.text
}

func generateError(t int) error {
	if t == 1 {
		return Error1{text: "generateError 1"}
	}

	return Error2{text: "generateError 2"}

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func wrapError() error {
	err1 := generateError(1)

	return fmt.Errorf("This is wrapping error: %w", err1)

}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	// Process the file here
	return nil
}

func main() {

	_, err := divide(1, 0)

	if err != nil {
		fmt.Println("This is error", err)
	}

	err = processFile("This_is_file_name_test")

	if err != nil {
		fmt.Println(err)
	}

	err = generateError(1)

	fmt.Printf("Is the error 1? %v", errors.Is(err, Error1{"generateError 1"}))
	fmt.Printf("Is the error 2? %v", errors.Is(err, Error2{"generateError 1"}))

	fmt.Println()

	var pErr Error1

	isSucces := errors.As(err, &pErr)

	fmt.Print(pErr, isSucces)

	werr := wrapError()

	fmt.Println()

	fmt.Println(errors.Is(werr, Error1{text: "generateError 1"}))

	var wErr12 Error1
	ok := errors.As(werr, &wErr12)

	fmt.Println(wErr12, ok)
	fmt.Println(io.EOF)

}
