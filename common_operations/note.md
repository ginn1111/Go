# Common operations
--- 
## 1. About some descriptions in operator explanations

## 2. Constant expressions

## 3. Arithmetic operators
  - About overflows: overflows allowed for non-constant and untyped constant, the opponent in the typed constant
  - About results of arithmetic operator operations:
    + Operations except bitwise
    + Bitwise operations 
  - About integer division and remainder operations
    ```go
      var a, b = 1.0, 0.0
      println(a/b, b/b) // +Inf NaN

      _ = int(a)/int(b) // compiles okay but panics at run time.

      // The following two lines fail to compile.
      println(1.0/0.0) // error: division by zero
      println(0.0/0.0) // error: division by zero
    ```
  - Using `op=` for binary arithmetic operations
  - The increment `++` and decrement `--` operators
    > Do not return so can't use as an expression
    ```go
      var a = 1
      println(a++) // error here
    ```

## 4. String concatenation operations

## 5. Boolean operations

## 6. Comparison operations 

## 7. More about constant expression  
  ```go
    package main

    const x = 3/2*0.1 // 3 and 2 is viewed as type int
    const y = 0.1*3/2 // 3 and 2 is viewed as type floating-point

    func main() {
      println(x) // +1.000000e-001
      println(y) // +1.500000e-001
    }
  ```

## 8. More operators
