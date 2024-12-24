# Basic control flows
---

## 1. An introduction of basic control flows in Go

## 2. `if-else` control flows 

## 3. `for` loop control flows
```go
  for InitSimpleStatement; Condition; PostSimpleStatement {
    // do something
  }
```
  - `PostSimpleStatement` must not be short variable declaration.
  - Can use as `while` when use using `condition-only` for loop form

## 4. `for-range` control flows
  - Can to loop the integer since Go version `1.22+`

## 5. `switch-case` control flows
```go
  switch InitSimpleStatement; CompareOperand0 {
  case CompareOperandList1:
    // do something
  case CompareOperandList2:
    // do something
  ...
  case CompareOperandListN:
    // do something
  default:
    // do something
  }
```
  - `InitSimpleStatement` can absent.
  - `CompareOperand0` is the typed boolean value. (can be untyped `nil`)
  - Auto break after each case expression
  - Use `fallthrough` statement to fall through to the next case
  - `fallthrough` must be the last statement in the case block code and do not within the `default` case

## 6. `goto` statement and label declaration

## 7. `break` and `continue` statement with label
  - The label must be defined before `break` or `continue`
  - Can break the breakable control flow although is not the innermost code block.


