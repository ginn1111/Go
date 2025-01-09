package main

import (
	"reflect"
	"unsafe"
)

func Float64Bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64FromBits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

func main() {
	s := "Go"
	slice := unsafe.Slice(unsafe.StringData(s), len(s))

	hdrSlice := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	hdrString := (*reflect.StringHeader)(unsafe.Pointer(&s))

	hdrSlice.Data = hdrString.Data
	hdrSlice.Len = hdrSlice.Len

}
