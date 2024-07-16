## Structs

A struct is a composite data type that groups together variables under a single name. Each variable within a struct is called a field. Structs allow you to create complex data types that represent real-world entities by combining different types of data.

### Defining a struct

To define a struct, use the type keyword followed by the name of the struct and the struct keyword:

```go
type person struct{
    name string
    age int
    email string
}
```

### Creating and Initializing a Struct

```go
p := person{
    name    :   "xyz",
    age     :   12
    Email   :   "xyz@gmail.com"
}
```

### Accessing Struct Fields

```go
fmt.Println(p.name)
```

## Methods

A method is a function with a special receiver argument. The receiver is a value or a pointer to a value of a specific type. Methods allow you to define behaviors associated with a particular type.

### Defining a Method

To define a method, specify the receiver in parentheses between the func keyword and the method name:

```go
func (p person) Greet(){
    fmt.Println("hello world")
}
```

### Calling a method

```go
p.Greet()
```

### Pointer Receivers

Methods can also have pointer receivers. Using a pointer receiver allows the method to modify the receiver's value:

```go
func (p *person) CelebrateBirthday() {
    p.Age++
}
```

In this example, CelebrateBirthday is a method with a pointer receiver. It increments the Age field of the Person struct.
