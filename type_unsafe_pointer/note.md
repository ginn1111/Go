# Type unsafe pointer

---

## 1. About the `unsafe` standard package

- There are 3 functions are provided in `unsafe` package:

  - `func Alignof(selector ArbitraryType) unintptr`
  - `func Offsetof(selector ArbitraryType) unintptr`
  - `func Sizeof(selector ArbitraryType) unintptr`

    > For the second function, if the selector denote embedded type, the field must be reachable without implicit pointer indirection.
    > These functions run at compile time.

    ```go
    type T struct {
      a int
    }

    type TT struct {
      *T
    }

    // fail to compiles
    fmt.Println(unsafe.Offsetof(TT.a))

    // its ok
    fmt.Println(unsafe.Offsetof(TT.T.a)) // the offset from address of type struct T not TT
    ```

  - There are more functions are more dangerous:
    - `func Add(ptr Pointer, len IntegerType) Pointer`
    - `func String(ptr *byte, len IntegerType) string`
    - `func StringData(s String) *byte`
    - `func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType`
    - `func SliceData(s []ArbitraryType) *ArbitraryType`

## 2. Unsafe pointer related conversion rules

- Pointer can explicit converted to unsafe pointer and vice versa.
- `uintptr` can explicit converted to unsafe pointer and vice versa.
  > Pointer shouldn't convert to `uintptr` and back with arithmetic.

## 3. Some facts in Go should we know

1. `unsafe.Pointer` is `pointer` and `uintptr` is `integer`
2. Unused memory block may be collected at any time
   > When convert unsafe pointer to `uintptr`
3. The address of some values might change at runtime
   > When the size of stack on the goroutine change, the addresses of value hosted on will change.
4. The life range of a value at runtime may be not at large as it looks in code
   > When convert unsafe pointer to `uintptr`
5. `*unsafe.Pointer` is a general safe pointer type
   > The safe pointer type with base type is `unsafe.Pointer`

## 4. How to use unsafe.Pointer type correctly

1. Pattern 1: convert `*T1` value to unsafe pointer, then convert unsafe pointer value to `*T2`.

- Use `unsafe.Pointer` like a intermediate to convert arbitrary type to each other
  > Size of T do not smaller than T2

```go

func Float64Bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64FromBits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}
```

2. Pattern 2: convert unsafe pointers to `uintptr`, then use the `uintptr` value. (not much useful)
3. Pattern 3: convert unsafe pointers to `uintptr`, do arithmetic with the `uintptr` value, then convert it back.
   > It is dangerous, because the `uintptr` value which represent the memory block may be collected, invalid or reallocated
4. Pattern 4: convert unsafe pointers to `uintpr` values as arguments of `syscall.Syscall` call.
   > The internal function of `syscall` prevent the memory block, whose referenced by do not be collected by the garbage collection
5. Pattern 5: convert the `uintpr` result of `reflect.Value.Pointer` or `reflect.value.UnsafeAddr` method call to unsafe pointers.

6. Pattern 6: convert a `reflect.SliceHeader.Data` or `reflect.StringHeader.Data` to unsafe pointers and vice versa.

- Just use these functions with existed `string` or `slice`. Do not allocate new memory block for these.
  - It is deprecated, using `unsafe.String`, `unsafe.StringData`, `unsafe.Slice` and `unsafe.SliceData` instead
