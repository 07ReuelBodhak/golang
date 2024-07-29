# Select Statement

the select statement is used to handle multiple channel operations. it blocks until one of its cases can proceed, then it executes that case. Heres the syntax:

```go
select {
case val := <-ch1:
    // code to execute when receiving from ch1
case ch2 <- 42:
    // code to execute when sending to ch2
default:
    // code to execute if no other case is ready
}
```

## Example Usage

### Basic Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		time.Sleep(3 * time.Second)
		ch1 <- "message 1"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		ch2 <- "message 2"
	}()

	for i := 0; i<2; i++{
		select{
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <- ch2:
			fmt.Println(msg2)
	}
}
}
```
