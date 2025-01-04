# Scheduler
---
## 1. OS Scheduler
  - `Thread` account for instructions execute sequentially
  - `Process` is created when start a program. Each `Process` has a initial `Thread`. The scheduling decision at `thread-level` do not `process`
  - `Thread` can run on `concurrency` or `parallel` mode. Each `Thread` as some props like `state`
  - The OS scheduler priority of threads, make the threads run as more as possible (reduce the idle time of CPU). The priority is not mean the `lower thread` will execution time's starved. The OS scheduler also responsible for minimizing latency time of each thread.

## 2. Executing instructions
  - The `Program Counter (PC)` or `Instruction Pointer (IP)` keep track of the next instruction will execute.

## 3. Thread state
  - There are three state of thread: `Runable` (want core to execute), `Waiting` (wait for io, network, ...) and `Executing` (running)

## 4. Types of work
  - There are two types of work: `CPU-bound` - the thread not do need switch to waiting state, otherwise, `IO-bound` always switch thread to waiting state.

## 5. Context switching
  - The performance nightmare of `CPU-bound`, because it make the `CPU` swap the runable thread and the executing thread.
  - Otherwise the second type of work reduce the idle time of `CPU`

## 6. Less is more
  - If the less threads are runable, the other threads will get more overtime. 
  - If the more threads are runable, the waste of time for `context switching`.

## 7. Find the balance
  - Using `Thread pool`, what is the min thread and max thread per core

## 8. Cache lines
 - False sharing in multithread applications.

## 9. Schedule decision scenario

## 10. Some note
  - Goroutine architecture more complex.
  - This work depend on M:P:G model. `P` do not larger than the number of processors in the computer (tip: `runtime.NumCPU()`)
  - `M` (is assigned to OS thread) is responsible for run the go code in the `goroutine` (G), which was attached on the `P`
  - The groroutine scheduler cooperating with OS scheduler.
  - The context-switching, asynchronous/synchronous syscall be handled by goroutine scheduler
  - Although, the context-switching, ... of goroutine work like OS thread, but in the deeper, the OS Thread do not move to waitting state, this mean goroutine make the `IO-bound` work in the `CPU-bound`, reduce the waste execute operations
  - Keywords: work-stealing, local run queue (LRQ), global run queue (GRQ), 1/61, network poller, syscall
