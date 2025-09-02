package myio

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readByLine(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// open and read file
func ex1() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := make([]byte, 100)
	_ = data

	readByLine(f)

	// count, err := f.Read(data)
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("read %d bytes: %q", count, data[:count])
}

// open and write file
func ex2() {
	file, err := os.OpenFile("data2.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := []byte("This is write data")

	count, err := file.Write(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Write the data with count %d", count)
}

// open and append file
func ex3() {
	file, err := os.OpenFile("data2.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := []byte("This is write data")

	count, err := file.Write(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Write the data with count %d", count)
}

func MYIO() {
	// ex1()
	// ex2()
	ex3()
}
