package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println(runtime.NumCPU())
}
