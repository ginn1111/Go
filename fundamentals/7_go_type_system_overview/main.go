package main

// The following new defined and source types are all
// basic types. The source types are all predeclared.
type (
	MyInt int
	Age   int
	Text  string
)

// The following new defined and source types are all
// composite types. The source types are all unnamed.
type IntPtr *int
type Book struct {
	author, title string
	pages         int
}
type Convert func(in0 int, in1 bool) (out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

type PersonAge map[string]int
type MessageQueue chan string
type Reader interface{ Read([]byte) int }

func f() {
	// The names of the three defined types
	// can be only used within the function.
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{ Read([]byte) int }
}
