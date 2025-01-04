# Context

---

## 1. The main usage

- Cancellation propagation and send value.

## 2. Type of context

- Background context: do not react any thing. (root context)
- With context, using two types:
  + There are some functions work with `cancellation`:
    + Creating a root context: `Context.Background()`
    + Creating a cancellation context: `Context.WithCancel(context.Context)`
    + Creating a cancellation context with timeout: `Context.WithTimeout(context.Context, time.Duration)`
  + There are some functions work with `value`:
    + Attaching the value to the context: `context.WithValue(context.Context, key any, value any)`
    + Get context's value from key: `context.Context.Value(key any)` 
    > Using `type assertion`  to get the correct key type

## 3. Some usage of context
  - Using in `http server`:
    + Cancellation:
      + The context will be attached into  `request`.
        > The package auto handle this `context` via cross server (from client to server)
      + If the client `cancel` the context via above methods, the `server` must implementation (using channel) to receive the `cancel signal`.
    + Value:
      > The package do not pass context's value cross server. Must implementation manually
      + Just using in `request's scoped`, do not using to send the data from the client to the server.
