# Pointers in Go
---

## 1. Memory address
  - The place where the value be stored
  - Exclude: `constant`, `function call` and explicit conversation

## 2. Value address
  - The address of value is the `start` address of memory segment

## 3. What are pointers
  - The value of pointer type is memory address 

## 4. Go pointer type and value
  - Unnamed `*T` pointer type, where `T` is the base type
  - If the `*T` is underlying type of named pointer, `T` be called base type of named pointer

## 5. About the word `reference`

## 6. How to get a pointer value and what are addressable value?
  - Using the keyword `new(T)` to allocate memory for type `T`, the return is a pointer type `*T`
  - Using `&` to take the address of addressable value. `&t` return the memory address of type `t` (hex string)

## 7. Pointer dereference
  - Using `*` to take the value of pointer type `T` with variable `t`. We call `*t` the value type `T` 
  - `*nil` cause the runtime panic

## 8. Why we do need a pointer?

## 9. Return the local pointer 

## 10. Restriction on pointers in go 
  - Don't support arithmetic operation like `p++` or `p--`, but `*p++` and `*p--` are valid
  - A pointer value can't be converted to an arbitrary type
    > The condition to convert:
    >  1. Identical underlying type
    >  2. Both is the unnamed pointer type and the base type of underlying type is identical
  - A pointer value can't be compared with values of arbitrary pointer type
    > The condition to compare:
    > 1. The types of the two Go pointers are identical
    > 2. The underlying type is identical and either of the two types is unnamed type
    > 3. One of these is `nil`
  - A pointer value can't be assigned to a value of arbitrary pointer type 

## 11. It's possible to break the pointer in go



