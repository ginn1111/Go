# Channel
---

## 1. Channel introduction
  - There are two ways to concurrent synchronization: `communitating by sharing memory` and `sharing memory by communitating`
  - Channel is a 2nd concurrent synchronization mechanism of above
  - The first way, which perform like mutex lock, ... in the `sync`  or `sync/atomic` standard package

## 2. Channel type and value
  - Channel types can be a bidirectional (`chan int`), send-only (`chan<- int`) or receive-only (`<-chan int`) channel
  - The two last types can convert to the first one, but not vice versa. And both of them can not convert to each other.
  - `int` is element type of channel
  - The zero-value of channel is `nil`, the value of channel must be created by `make(chan int, [capacity])`.
    + If the `capacity` is 0, the channel is `unbuffered`, otherwise the channel is `buffered`

## 3. Channel comparison
  - All channel values are comparable
  - Multiple-parts, so when two channels share underlying parts -> the equal comparison (==) is `true`

## 4. Channel operations
  - `close` channel (not is the receive-only channel)
  - Send value to channel `ch <- value`
  - Receive value from channel `value, isSentBefore := <- ch`
  - `cap(ch)`; `cap(ch(nil)) = 0`, capacity of channel
  - `len(ch)`; `len(ch(nil)) = 0`, length of channel, the elements do not taken out
  > These operations are not synchronized

## 5. Detailed explanations for channel operations
  - There are three queues exist in internal structure of channel:
    + Receiving goroutine queue
    + Value buffer queue
    + Sending goroutine queue
    + Three queues work and associate with each other .
    + If the receiving goroutine queue is empty, the value buffer queue must be not empty.
    + If the value buffer queue is not full, the sending goroutine queue must be empty.
    + If the channel is unbuffered, most the sending or receiving goroutine queue is empty 
    + If the non-nil non-closed channel close, both receiving and sending goroutine queues must be empty

## 6. Some channel use cases example

## 7. Channel element value are transferred by copied
  - Like assignment or arguments passing, transferred values are copied (just the direct part with multiple-parts value ) 
  - The limitation of element type's size of the base type channel is 2^16 (65536)
  - In this circumstances, the element type should be `pointer` type

## 8. About channel and goroutine garbage collection
  - The garbage collection of goroutine if it is exited.
  - The garbage collection of channel if it is not referenced by any goroutine.

## 9. Channel send and receive operations are simple statements

## 10. `for-range` on channel
  - The `for-range` can be used for get all the received value from channel or get the values until the value buffer queue is empty.
  - When using the `for-range` with channel, just use one iteration variable: `for v := change aChannel`; using `for-while-style` to get the second result of receive channel operation: 
  ```go
  for {
    v, ok := <- aChannel;
    if !ok {
      break;
    }
  }

  ```

## 11. `select-case` control flow code blocks
  - The case operation must be send or receive channel operation. If both cases is in `blocking` state, the `default` case (if present) will be selected.
  - The case operations in `select-case` are selected randomly.

## 12. The implementation of the select mechanism

## 13. More
