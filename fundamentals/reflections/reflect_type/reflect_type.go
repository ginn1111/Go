package reflect_type

import (
	. "fmt"
	"reflect"
)

func WithInterface() {
	type T interface {
		M(int) int
		m(int) bool
	}

	type TS []interface {
		M(int) int
	}

	tt := reflect.TypeOf(new(T))
	ts := reflect.TypeOf(TS{})

	Println("Kind:", tt.Kind(), ts.Kind())

	ttE := tt.Elem()
	tsE := ts.Elem()

	Println("Element Type:", ttE.Kind(), tsE.Kind())

	Println("Implementation: ", ttE.Implements(tsE)) // true
	Println("Implementation: ", tsE.Implements(ttE)) // false
}

type T struct {
	f string `max:"99" min:"0"`
	A string
}

func (t T) M() {}
func (t T) m() {}

func WithNonInterface() {
	type A = [16]int16
	type B = [16]func() // function, slice, map, struct or array with incomparable element type
	var c chan map[A][]byte

	ta := reflect.TypeOf(A{})
	tc := reflect.TypeOf(c)
	tb := reflect.TypeOf(B{})

	Println("Kind:", ta.Kind(), tc.Kind())

	Println("Channel direction:", tc.ChanDir())

	Println("Reflection with key of map:", tc.Elem().Key().Elem().Bits())

	Println("Is comparable:", ta.Comparable()) // true
	Println("Is comparable:", tb.Comparable()) // false
	Println("Is comparable:", tc.Comparable()) // true

	tt := reflect.TypeOf(T{})

	Println("Number of methods and fields:", tt.NumMethod(), tt.NumField())

	tm, okm := tt.MethodByName("M")
	tf, _ := tt.FieldByName("f")

	Println("Method by name", tm, okm)
	Println("Field by name", tf)

	tagV, okTag := tf.Tag.Lookup("min")

	Println("Struct tag:", tf.Tag.Get("max"), tagV, okTag)

}
