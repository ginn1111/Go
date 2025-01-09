package main

import (
	"fmt"
	"unicode/utf8"
)

func RunesToBytes(runes []rune) []byte {
	n := 0

	for _, r := range runes {
		n += utf8.RuneLen(r)
	}

	n, bs := 0, make([]byte, n)

	for _, r := range runes {
		n += utf8.EncodeRune(bs[n:], r)
	}

	return bs
}

func main() {

	hello := []byte("Hello")
	world := "world"

	fmt.Println(string(append(hello, world...)))

	helloWorld := make([]byte, len(hello)+len(world))

	copy(helloWorld, hello)
	copy(helloWorld[len(hello):], world)

	fmt.Println(string(helloWorld))

}
