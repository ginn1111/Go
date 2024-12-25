# Go type system overview
---
## 1. Concept: basic types

## 2. Concept: composite types
  - pointer, struct, function, container, channel, interface
  - Unamed composite types
  ```go
  // Assume T is an arbitrary type and Tkey is
  // a type supporting comparison (== and !=).

  *T         // a pointer type
  [5]T       // an array type
  []T        // a slice type
  map[Tkey]T // a map type

  // a struct type
  struct {
    name string
    age  int
  }

  // a function type
  func(int) (bool, string)

  // an interface type
  interface {
    Method0(string) int
    Method1() (int, bool)
  }

  // some channel types
  chan T
  chan<- T
  <-chan T
```
## 3. Fact: kind of types: 26

## 4. Syntax: type definition
  - Type definition declarations
  ```go
  // Define a solo new type.
  type NewTypeName SourceType

  // Define multiple new types together.
  type (
    NewTypeName1 SourceType1
    NewTypeName2 SourceType2
  )
```
  - The defined type and the source type must share the underlying type, and the values and convert to each other.
                                                                 
## 5. Concept: custom generic types and instantiated types
  - When using the generic type, it must be instantiated type.
  - Constraints and type parameters

## 6. Concept: named type and unnamed typed
  - Named types:
    + The predeclared types.
    + a defined type.
    + a instantiated type.
    + a type parameter type.
  - Unnamed types: composite types 

## 7. Syntax: type alias declarations

## 8. Concept: underlying type
  - For built-int type, it is itself.
  - The `pointer` type defined in the `unsafe` package is itself.
  - The underlying type of the unnamed type is composite type, is itself.
  - In the type definition, the source type and the definition type must the same underlying type.
  - The underlying be tracked to the `unnamed` type (composite types literal or built-in types)

## 9. Concept: values
  - Be called the instance of type
  - There are some kind of values: `literal`, `named constants`, `variables` and `expressions`. The other is `function literal` and `compsite literal`
  - The function declaration consist of: `function literal` (function value) and `identifier` (the name)
  - The composite literal are used to represent the values of struct type or container type.
  - There are no value literal type for `channel`, `pointer` and `interface`

## 10. Concept: value parts
  - The values are stored in the memory, include the direct and indirect parts.
  - The indirect underlying parts of value be referenced by it direct via (`pointer`, `unsafe.Pointer`).

## 11. Concept: values size
  - The bytes occupies by the direct part of the value

## 12. Concept: base type and pointer type
  - Pointer type `*T` where `T` is the base type

## 13. Concept: fields of struct type

## 14. Concept: signature of function type
  - The signature of function consist of the parameters definition list and out result definition list.

## 15. Concept: method and method set
  - Method can be called `member function`.
  - Method set of a type is composed of all the methods of the type.

## 16. Concept: dynamic type and dynamic value of an interface type
  - Can boxing the non-interface type
  - The value is boxed the interface value be called `dynamic value`.
  - The type is boxed the interface type be called `dynamic type`. 
  - The interface boxing nothing - the zero value of interface type (do not have dynamic type or value).
  - The method set of interface type of non-interface is wrapped the set of interface type. We say the type `implements` the interface type.




