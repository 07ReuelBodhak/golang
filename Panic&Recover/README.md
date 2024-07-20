# Panic and Recover

### `panic`

The `panic` function is used to stop the normal execution of a Go program and begin panicking. When a function calls `panic`, the program starts to unwind the stack, running any deferred functions as it goes. This process continues until the program terminates or a `recover` function is encountered.

Here's an example of using `panic`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Start of main")
    panic("something went wrong")
    fmt.Println("End of main") // This line will never be executed
}
```

In this example, the message "something went wrong" will be displayed, and the program will terminate after executing any deferred functions in the call stack.

### `recover`

The `recover` function is used to regain control of a panicking goroutine. `recover` is only useful when called inside a deferred function. If `recover` is called inside a deferred function and the goroutine is panicking, `recover` will stop the panic and return the value that was passed to `panic`. If the goroutine is not panicking, `recover` will return `nil`.

Here's an example of using `recover`:

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Start of main")
    panic("something went wrong")
    fmt.Println("End of main") // This line will never be executed
}
```

In this example, the deferred function recovers from the panic, so the program does not terminate. Instead, it prints the message "Recovered from panic: something went wrong" and continues execution.

### Detailed Explanation and Use Cases

#### How `panic` Works

- When `panic` is called, it stops the normal execution of the program.
- The runtime starts unwinding the stack, executing any deferred functions along the way.
- If no `recover` is called, the program prints the panic message and terminates.

#### How `recover` Works

- `recover` can only be used inside deferred functions.
- When called, `recover` stops the panic and returns the value passed to `panic`.
- If there is no ongoing panic, `recover` returns `nil`.

#### Use Cases for `panic` and `recover`

- **Handling Unexpected Errors**: `panic` can be used to handle unexpected conditions that the program cannot recover from, such as accessing an invalid memory address or running out of resources.
- **Defensive Programming**: `recover` can be used to create robust programs that can gracefully handle panics and continue execution or perform cleanup operations.
- **Libraries and Packages**: In library code, `panic` can be used for internal errors, while `recover` allows the library to avoid crashing the entire program using it.
