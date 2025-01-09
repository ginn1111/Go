# String
---

## 1. Some simple facts of string
  - `s` type string is `constant`, `start` and `end` are `constant` but `s[start:end]` is NOT the constant

## 2. String encoding and Unicode code point
  - The unit of Unicode is `code point`
  - `code point` are represented as `rune`
  - `code point` be stored at one or several bytes (up to four bytes - `rune` - `int32`)

## 3. String related conversion
  - `int` types can explicit convert to `string` but not vice versa
  - `string` can explicit convert to `byte` slice and vice versa
  - `string` can explicit convert to `rune` slice and vice versa
  - Go do not support built-in function to convert `rune` slice to `byte` slice. From `byte` slice to `rune` slice, use `bytes.Runes([]byte)` to convert `byte` slice to `rune` slice
  - To convert `rune` to `byte`, use `unice/utf8` standard package (`uft8.RuneLen(rune)` and `utf8.EncodeRune([]byte, rune)`)

## 4. Compiler optimization for conversion between strings and bytes
  - Some cases, the compiler do not duplicate copy the underlying elements of string and byte:
    + `string` to `byte slice` follow by `range`
    + `byte slice` to `string` when use at map key in map element retrieval indexing
    + `byte slice` to `string` in comparison
    + `byte slice` to `string` when concatenation with at least one `non-blank string constant`

## 5. `for-range` on string
  - Iterate the per `rune` not `byte`

## 6. More string concatenation methods
  - `+` is the best way!!!

## 7. Sugar: Use strings as byte slice
  - When using `copy` and `append` to copy and append byte slice, `string` can pass to the argument as `byte slice` type
  ```go
  hello := []byte("Hello")
  world := "world"

  append(hello, world...)

  helloWorld := make([]byte, len(hello) + len(world))

  copy(helloWorld, hello)
  copy(helloWorld[len(hello):], world)

```
## 8. More about string comparison
  - Avoid comparison two long strings when both do not share the underlying elements.
