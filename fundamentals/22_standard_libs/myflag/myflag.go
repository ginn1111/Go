package myflag

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func MyFlag() {

	var userColor bool

	flag.BoolVar(&userColor, "color", false, "display colorized output")

	flag.Parse()

	if userColor {
		colorize(ColorBlue, "Hello flag package")
		return
	}

	fmt.Println("Hello flag package")

}

func CloneHead() {
	var n int

	flag.IntVar(&n, "n", 5, "The lines will be read")

	flag.Parse()

	var in io.Reader

	if filename := flag.Arg(0); filename != "" {
		file, err := os.OpenFile(filename, os.O_RDONLY, 0600)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		in = file

	} else {
		in = os.Stdin
	}

	scanner := bufio.NewScanner(in)

	for range n {
		if !scanner.Scan() {
			break
		}

		fmt.Println(scanner.Text())
	}
}
