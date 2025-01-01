# Method in go
---

## 1. Method declaration
  - Declaring a method for  type `T` or `*T` when satisfy the following conditions:
    + `T` is a defined type. 
    + `T` must be defined in the same package as the method declaration.
    + `T` must not be a pointer type.
    + `T` must not be an interface type.
    + `T` and `*T` are called receiver of methods.
    + `*T` be called `pointer receiver`. `T` be called `value receiver`

## 2. Each method corresponds to an implicit function
  - When declare a method, the compiler will implicit declare the correspond function for it.

## 3. Implicit method with pointer receiver
  - When declare a method for receiver value type, the compiler will implicit declare two functions, one for value receiver and one for correspond pointer receiver

## 4. Method specifications and method sets
  - The method specification is a method declaration remove the `func` keyword.
  - The method sets is all of method specification of a type
  - The method sets of type `T` is always subsets of method sets of type `*T` 
  - The are some type, whose the method set is blank:
    + The `builtin` type
    + The defined pointer type
    + The pointer type whose base type is interface type or pointer type
    + `unnamed` slice, array, map, function and channel types
                                       
## 5. Method values and method call
  - The typed `nil` value also has member function if its type

## 6. Receiver arguments pass by copied

## 7. Method value normalization

## 8. Method value evaluation
  - The receiver value will be evaluated when the `v.m` is evaluated.
  - For method declaration for `value receiver`:
    + `v1 := v`; `v1.m` is normalize already, the value of `v` will copy and save at this time.
    + `v2 := &v` `v2.m -> (*v2).m (normalization)`, the value is the address of `v` so the modify of `v` will reflect to `v2` 
  - For method declaration for `pointer receiver`:
    + `v1 := v`; `v1.m ->  (&v1).m (normalization)`, the value of `v1` is address of `v` so the modify of `v` will reflect to `v1`
    + `v2 := &v`; `v2.m` is normalize already, the value of `v2` is address of `v` so the modify of `v` will reflect to `v2`

  > The receiver value of the method of pointer receiver will reflect when the deference of pointer is modified.
  > Otherwise, the receiver value of the method of value receiver just reflect when the value of receiver is the address 

## 9. The defined type doesn't obtain the methods declared explicitly for the source type use in its definition

## 10. Should the method will be declared with pointer receiver or value receiver
- If it is hard to make a decision whether a method should use a pointer receiver or a value receiver, then just choose the pointer receiver way.



