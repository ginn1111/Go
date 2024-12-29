# Array, Slice and Map
---

## 1. Simple overview of container types and values
  - Array: one direct part the continuous memory be hosted in the direct part
  - Slice: one direct part and underlying parts, the continuous memory be hosted in the underlying parts
  - Map: one direct part and underlying parts.

## 2. Literal representation of Unnamed container types
  - Array: `[N]T` 
  - Slice: `[]T` 
  - Array: `map[K]T` 

  - Where: `N` is non-negative integer, `T` is an arbitrary type, `K` is comparable type

## 3. Container value literal
  - Array: can using value literal like this: `[...]T{1,2,3}` to the compiler deduce the length of the array 
  - The element indexed key are optional when using container value literal of `array` and `slice` type. If we use the indexes, will be non-negative and constant
  - If the indexes is absent, it will use the value of previous index + 1. The 0 index is `0`

## 4. Literal Representation of zero value of container types
  - Array: `[100]int{}`, every elements are zero type. In particular this case, is `0` 
  - Slice and Map: `nil`
  - `T[]{}`, `T[](nil)`, `map[K]T{}` and `map[K]T(nil)`

## 5. Composite literals are unaddressable but can be taken the addresses

## 6. Nested composite literals can be simplified

```go
  type Language struct {
    Name string
    Version int
  }

 map1 := map[Language]map[string]int{
  {"Go", 1}: {
    "a": 1,
    "b": 2,
    }
  }
```
## 7. Compare container value
  - Map and slice are incomparable values
  - But the array can be comparable when the element types is comparable (How? The same as  comparable `struct` )

## 8. Check length and capacity of container values
  - Using `len` or `cap` built-in functions to retrieve the length and capacity of container values
  - Array: length and capacity are the same
  - Map: length is the number of key-value pairs, capacity is unlimited 
  - Slice: length and capacity can be difference 

## 9. Retrieve and modify container elements
  - For `slice` and `array`, the index is non-negative and less than the length, otherwise the panic will occurs
  - For `map`, the key type must be comparable, otherwise the panic will occurs. 
    + `v[k] = x` where `v`  is `nil` value, the panic will occurs when using `v[k]` as the destination value. 
    + But when using `value, isPresent := v[k]` the panic will not occurs even if `v` is `nil` and the second value represent where or not the value is absent.

## 10. Container assignment
  - Array: copy all the elements to the new array
  - Slice and Map will just copy the direct part (its mean the underlying parts will be shared)

## 11. Append and delete container elements
  - Array: can not append or delete the elements of its, just modify
  - Map: append like modify. Delete by using `delete(m,k)` built-in function (`delete` do not throw a panic if the key is not present or even map is `nil`)
  - Slice: is more complex than two above container types for append and delete: for append elements
    + Depend on the capacity of slice, whether or not results slice and base slice is share some elements
    > - If the number of elements be append to base slice larger than the redundant slots (length - capacity) -> the results slice and the
    >   base slice do not share any elements. In the practice, we usually assign the results slice to the base slice
    > - Otherwise, the results slice and base slice share some elements (just the base slice's elements, because the length is over of the result).
    + For delete operation, we always use the `append` built-in function to perform it.
    
## 12. Create slice and map with the built-in `make` function
  - Array: don't use the `make`
  - Map: `make(M, capacity)` or `make(M)`, capacity is non-position will be ignored
  - Slice: `make(S, length)` or `make(S, length, capacity)`, the elements value of slice is zero-value

## 13. Allocate containers with the built-in `new` function

## 14. Addressability of container
  - Array: the elements of addressable value array also are addressable, the same as unaddressable array values
  - Map: unaddressable
  - Slice: always addressable

## 15. Derive slice from array and slice
  - The derived slice and the base slice or addressable array are share some elements.
  - `s[low:high:max]`, the derived slice will has length = `high - low`, capacity = `max - low` 
  - `low` can be omitted when low = 0, `high` can be omitted when high = len(s), if `max` is omitted, max = cap(s)
  - subslice of `nil` slice is a `nil` slice

## 16. Convert slice to array pointer
  ```go
  s := make([]int, 10)
  a := (*[10]int)(s) // its ok
  a2 := (*[9]int)(s) // its ok
  a3 := (*[11]int)(s) // panic occurs
```

## 17. Convert slice to array
  - The same as convert slice to array pointer
  - Do not share any elements

## 18. Copy slice with the `copy` built-in function
  - Copy the slice or array to the another one, where the both share the underlying element type
  - Return the number of value to copied: `numsCopied := copy([]D, []S)`

## 19. Container elements iterations
  - Array: copy all the elements to the anonymous array
  - Slice and map just copy the direct part
  - BTW, the element of container types will copy of each iteration
  - Iterator over `nil` map or slice is allowed
  - The instances key-pair of `for-range` is difference between prior Go.1.21 and Go.1.22
 
## 20. Use array pointers as arrays
  - subslice of `nil` pointer array will cause a panic

## 21. The `memclr` optimization
  - when assign the element value to the zero value of the element type, the compiler will use the `memclr` optimization

## 22. The `clear` function
  - Clear value of map or reset value of slice
  - Clear the `NaN` key of map

## 23. Call to the built-in `len` or `cap` function may be evaluated at compile time
                                      .
## 24. Modify the length and capacity of a slice individually 
  - Using `refect`
               .
## 25. More slice manipulations
  - Clone: 
    + `append(s[:0:0], s...)`
    + `append([]T(nil), s...` if s is blank non-nil, the result is nil
    + `sClone := make([]T, len(s)); copy(sClone, s)`
    + `sClone := append(make([]T, 0, len(s)), s...)`
    + In practice: 
      ```go
      if s != nil {
        sClone = make([]T, len(s))
        copy(sClone, s)
      }
      ```
  - Delete: delete `from` to `to` of slice `s`
    + `append(s[:from], s[to:]...)`
    + `s[:from + copy(s[from:], s[to:])]`
    - ```go
      if n:= to - from; len(s) - to < n {
        copy(s[from:to], s[to:])
      } else {
        copy(s[from:to], s[len(s) - n:])
      }
      s = s[:len(s) - (from - to)]

    ```
  - Delete one element: `i`
    + `append(s[:i], s[i+1:]...)`
    + `s[:from+copy(s[i:], s[i+1:])]`
    + ```go
      s[i] = s[len(s) - 1]
      s = s[:len(s) - 1]

    ```
  - Delete slice elements conditionally
  - Insert all elements to the slice
    + `append(s[:i],append(elements, s[i:]...)...)`
    + ```go
      if cap(s) >= len(s) + len(elements) {
        s = s[:len(s)+len(elements)] // expand len equal capacity to copy
        copy(s[i+len(elements):], s[i:]) // copy(s[i+1:], s[i:])
        copy(s[i:], elements) // s[i] = a
      } else {
        x := make([]T, 0, len(s) + len(elements))
        x = append(x, s[:i]...)
        x = append(x, elements...)
        x = append(x, s[i:]...)
        s = x
      }
    ```
  - Pop front (shift): `s, e := s[1:], s[0]`
  - Pop back: `s, e = s[:len(s)-1], s[len(s)-1]`
  - push front: `s = append([]T{e}, s...)`
  - push back: `s = append(s, e)`
  - More operations

## 26. Use map to simulate set
  - `map[K]struct{}`

## 27. Container related operation are not synchronized internally  


