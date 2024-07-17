# Interface

Interfaces in Go are a powerful feature that allows you to define behavior without specifying how the behavior is implemented. They are a key part of Go's type system and enable polymorphism, making it easier to write flexible and reusable code.

```go
type InterfaceName interface{
    method1(param1,param2) ReturnType,
    method2() ReturnType
}
```

## Empty Interface

The empty interface interface{} is a special case in Go. It has zero methods and is satisfied by any type, making it similar to Object in other languages. It’s often used when the exact type of the value is not known ahead of time.

```go
func Describe(i interface{}){
    fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    Describe(42)         // Output: (42, int)
    Describe("hello")    // Output: (hello, string)
    Describe(true)       // Output: (true, bool)
}
```

## Type Assertions and Type Switches

When you have a value of an interface type, you can use a type assertion to extract its underlying concrete type. Here's how you do it:

```go
var i interface{} = "hello"

s, ok := i.(string)
if ok {
    fmt.Println(s)
}else{
    fmt.Println("i is not a string")
}
```

For multiple types, you can use a type switch:

```go
switch v := i.(type) {
case string:
    fmt.Println("i is a string:", v)
case int:
    fmt.Println("i is an int:", v)
default:
    fmt.Printf("i is of type %T\n", v)
}

```
