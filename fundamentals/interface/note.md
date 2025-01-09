# Interface

--

## 1. Interface type and interface set

- Interface type defines some type requirements. The types satisfy this requirements, these are called type sets of the interface.
- The requirements in the interface type are called `interface element`. Interface element is either method specification of type name.
- The type sets of interface type is infinite.
- `any` is predeclared alias for `interface{}` literal type.

## 2. Method sets of types

- For none-interface type, composed of all methods specification of the type (explicit or implicit)
- For interface type, composed of all methods specification either directly or indirectly through embedding other types.

## 3. Basic interface types

- The interface types which can used as value types that be called `basic interface type`. Otherwise, it's called `non-basic interface type` or `constraint-only interface type` (approximate type or union type)

## 4. Implementations

- If type `T` implement interface type `X`, then the method sets of type `T` is supper-set of method sets of interface type `X`. Generaly, not vice versa. But when `X` is basic interface type, then vice versa

## 5. Value boxing

- If a type `T` implement basic interface type `X`, then any value of type `T` will implicitly converted to the type `X`.
- In other words, value types `T` are assignable to value types `X`.
- If `T` is non-interface type, the complexity of assignment is `O(n)`, where `n` is size of type `T`.
- If `T` is other interface type, the complexity is `O(1)`.
- `Dynamic value` and `Dynamic type` (type of `dynamic value`), which be stored in the interface value.
- Any type implement the `blank interface type` (`interface{}`).
- When assign a value to `blank interface type`, firstly, it convert to its `default type`.
- In fact, the `Go compiler` built a table about implementation information of interface type and corresponding non-interface type, consist of:
  - The information about dynamic type.
  - The methods table of the interface type and non-interface type.

## 6. Polymorphism

- Thanks to implementation information of go compiler, When any types implement a interface type `X`, then the value type `X` can call the methods of corresponding types, even if the behave of this type is difference.

## 7. Reflection

- Type assertion
  - `i.(T)`, where:
    - `i` is a value, `T` either `interface type` or `non-interface type`
    - Calling type of `i` is `I`, the `type assertion` will success if:
      - `I` and `T` is an interface type, these types can `type assertion` to each other if `I` implement `T` or `T` implement `I`
      - `I` is an interface type, and `T` is not an interface type, `T` must implement interface type `I`
      - `I` is not an interface type, and `T` is an interface type, `I` must implement interface type `T`
      - `nil` interface value can not `type assertion` success, because it do not implement any interface types
- `type-switch` control flow block

  ```go
   var x  = []interface{}{
   struct{name string; age int}{name: "Thuan", age: 24},
     1,2,
     "Go",
   }

  for _, v := range x {
    switch i := v.(type) {
    case string:
      fmt.Println("String type", i)
    case struct {
      name string
      age  int
    }:

      fmt.Println("Struct type", i)
    default:
      fmt.Println("Default", i)

    }
  }
  ```

## 8. More about interface type in go

- Comparison involving interface value
  1. If one of value is `nil` interface value, then the comparison result whether or not the another also is `nil` interface value.
  2. If the dynamic type of both values is difference, then the comparison result is `false`.
  3. If the dynamic type is the same:
  - If the dynamic type of one is `incomparable`, then the `panic` will occur.
  - Otherwise, the comparison result whether or not both dynamic value are equal.
- The internal structure of interface values: The `non-blank` interface value has the method set of its dynamic type
- Pointer dynamic value vs non-pointer dynamic value
  - Avoid boxing large size value, boxing their pointer type instead.
- Value of `[]T` can't be directly converted to `[]I`, even if type `T` implements interface type `I`
- Each method specified in an interface type corresponds to an implicit function:

  ```go
    type I interface {
      m(int)bool
    }

    type T string

    func (t T) m(i int) bool {
      // do some thing
    }

    i := T("Hello")
    i.m(1) // (1)
    I.m(i, 1) // the same as (1)
    interface{m(int)bool}.m(i, 1) // the same as (1)
  ```
