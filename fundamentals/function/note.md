# Function
---

## 1. Function signature and function type
  - Function signature the same as function type.
  - Variables (parameters and results) name is not important.
  - `func (int, int) int`

## 2. Variadic parameters and variadic function type  
  - The function has variadic parameter are called variadic function.
  - `func vadiadicFunc(a int, b ...string)` -> at most one variadic parameter
  - The variadic parameter is slice type `[]T`

## 3. Function type is incomparable types
  - Can compare with untyped value `nil`

## 4. Function prototype
  - Composite of function signature and function name: `func GCD(int, int) int`
  - Function declaration is combined of function prototype and function body

## 5. Variadic function declaration and variadic function call
  - Variadic function declaration
    ```go
      func Sum(nums ...int) int {
        // body
      }
    ```
  - Variadic function call 
    ```go
      Sum(1, 2, 3) // (1)
      Sum([]int{1,2,3}...) // (2)
      Sum()

      // Sum(1, []int{1,2,3}) // compile fail
    ```
    + Do not mix the one and two ways each other

## 6. More about function declaration and call
  - Function whose name can be duplicated: the `init` function or `blank identifier`
  - Some function calls are evaluated at compile time: `len`, `cap`, ...
  – All function arguments pass by copy.
  - Function declaration without body, use in Go assembly
  - Some function with results do not required to return: infinite stuck functions

## 7. The results of custom function can be discarded, not true for all builtin functions (except `recover` and `copy`)

## 8. Use function calls as expressions
  - When the function results is a single value, it can be used as single expression.
  - Otherwise, the function results do not mixed with other expressions. 
    ```go
      func Test(a, b int) (int, int) {
        
      }

      func Test2(a ...int) (int, int) {
        
      }

      _ := Test2(Test(1,2)...) // ok
      _, _, _ := Test(1,2), 3 // compile fail
    ```
## 9. Function values
  - The function value is the same as function signature + function body (omit the function name)
  - When assign a function to another one, the both functions will share the underlying parts

## 10. Range over function
  ```go
  func Loop3(yield func(num, square int) bool) {
    for i := range 3 {
      if !yield(i, i*i) {
        return
      }
    }
  }

  func main() {
    for i, ii := range Loop3 {
      fmt.Println(i, ii)
    }
  }
```
