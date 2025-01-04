package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	buf := bytes.NewBuffer(make([]byte, 5, 10))
	_ = bytes.NewBufferString("Thuandz")

	fmt.Println("buf.Available (cap - len): ", buf.Available())
	fmt.Println("buf.AvailableBuffer: ", buf.AvailableBuffer())
	fmt.Println(buf.Cap())
	fmt.Println(buf.Len())

	_2Bytes := buf.Next(2)

	copy(_2Bytes, []byte("Th"))
	_2Bytesp := bytes.NewBuffer(_2Bytes)

	fmt.Println(_2Bytesp.Len(), buf.Len())
	_2Bytesp.Read(buf.Bytes())
	err := _2Bytesp.UnreadByte()
	fmt.Println(err)
	fmt.Println(_2Bytesp.Len(), buf.Len())
	fmt.Println(_2Bytesp.String(), buf.String())

	buf.WriteString("uandz")
	buf.WriteTo(os.Stdout)

}
