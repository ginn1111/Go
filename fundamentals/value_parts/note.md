# Value parts
---

## 1. Two categories of Go types
  - The value types be hosted on the single memory block 
  - The value types be hosted on the multiple memory blocks

## 2. Two kinds of pointer in Go
  - Pointer and unsafe Pointer
  - Pointer type, pointer wrapper type and pointer holder type
  - All pointer types and pointer wrapper types are pointer holder types

## 3. (Possible) Internal Definition of the Types in the Second Category

## 4. Underlying value parts are not copied in the value assignment
  - Value assignment: pure assignment, parameters passing
  - Just copy the direct part (its mean the underlying value parts will share to each other)

## 5. About reference type and reference value
  - Using pointer holder (pointer types and pointer wrapper types)
