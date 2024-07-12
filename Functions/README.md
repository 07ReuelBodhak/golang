# Functions in go

functions are block of code that are executed when it is called

## All types of function

1.  Basic function

    syntax :

    ```go
    func function_name(parameter type)return type{
        return value
    }
    ```

2.  multiple return value : this function can return multiple return value

    syntax :

    ```go
    func function_name(parameter1 type,parameter2 type)(return1 type,return2 type){
        return value1,value2
    }
    ```

3.  Named Return value : you can name the return values, which can act as a variable within the function

    Example:

    ```go
    func split(sum int) (x,y int){
        x = sum * 4 / 9
        y = sum - x
        return
    }
    ```

4.  Varriadic Function : Variadic function allows you yo pass a variable name of arguments

    syntax :

    ```go
    func function_name(parameter ...type) returnType{
        return
    }
    ```

5.  Annonymous function : Annonymous functions are function that has no name

    syntax :

    ```go
    func (parameter value) returnType{

    }
    ```

6.  Highorder function : Higher-order functions are a fundamental concept in functional programming and can be effectively used in Go, although Go is not primarily a functional language. A higher-order function is a function that either:

    1. Takes one or more functions as arguments, or

    2. Returns a function as its result.

    syntax :

    ```go
    func function_name(f func(type1) type2,parameter type) returntype{

    }
    ```

7.  Mthods : Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it. With the help of the receiver argument, the method can access the properties of the receiver.

    Syntax :

    ```go
    func (reciever_name type) methodName(parameters)(returnType){

    }
    ```

8.  Functions Literals : Function literals are anonymous functions defined inline and executed immediately.

    Syntax :

    ```go
    result := func(parameter1 value,parameter2 value) returnType{

    }(arguments1,arguments2)
    ```
