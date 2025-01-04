# Mutex
---
## 1. The usage
  - One of the techniques to work with `concurrency programming`. `Communicating by sharing memory`.
  - Using from `sync` standard package.

## 2. Some notes
  - The `zero value` of mutex if `Unlockl mutex`
  - There are two types of mutex: `sync.Mutex` and `sync.RWMutex`
  - The second one allow `read parallel` but `lock` for each `reads` and `writes`
  - We can bundle `sync.Mutex` in the `struct type`
    ```go
      type IntType struct{
        val int
        sync.Mutex
      }
    ```

## 3. Common pitfalls
  - Make sure that call `UnLock` after `Lock` function (tips: using `defer` function)
  -  Releasing the mutex as soon as possible (tips: using `IIFE`)
