# Function declaration and function call
---

## 1. Function declarations
```go
  func gcd(a, b int) int {
    // body
    if b == 0 {
      return a
    }

    return gcd(b, a%b) // anonymous result
  }

  func test(a, b int) (c, d int) {
    c = a + b // name result
    d = a - b // name result
    return
  }
```

## 2. Function call
  - Like JS, func in go can hoist declaration

## 3. Exiting (or Returning) phase of function call
  - The exiting phase start then the function call be returned

## 4. Anonymous function
  - Like declaration function but do not have the name
  ```go
  x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.
```
## 5. Built-in functions

## 6. More about functions


