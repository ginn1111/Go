# Goroutines, deferred function call and panic/recover
---
## 1. Goroutines
  - A function call with `go` keyword, the function call do not have any arguments and the results will discard.

## 2. Concurrency and Synchronization

  - Resolve the `data race` problem (two or more threads write or read the same memory block)

## 3. Goroutines state
  - There are two goroutine's state: `running` and `blocking`
  - If all the goroutine go to the `blocking` state -> `deadlock`
  - There are some subs-state of `running` state: `queuing`, `executing`

## 4. Deferred function call
  - The `defer` with the function call be called `defer statement`.
  - The arguments of `defer statement` will evaluate when the `defer`  function be pushed to the defer queue (first-int, last-out)

## 5. Panic and recover
  - The mechanism of go to handle exception
  - Some fatal errors do not recover: like `stack overflow` or `out of memory`

