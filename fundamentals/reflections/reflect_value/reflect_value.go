package reflect_value

import (
	"fmt"
	"reflect"
)

func WithInterface() {
	fmt.Println("WithInterface in reflectValue")
}
func WithNoneInterface() {
	n := 5
	p := &n

	vp := reflect.ValueOf(p)
	fmt.Println("Can set and can addressable: ", vp.CanSet(), vp.CanAddr())

	ve := vp.Elem()

	fmt.Println("Can set and can addressable of base type: ", ve.CanSet(), ve.CanAddr(), ve.Kind())

	ve.Set(reflect.ValueOf(123))

	fmt.Println("New value:", n)

}
