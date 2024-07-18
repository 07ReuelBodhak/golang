# Error Handling

## Basic Concept

1. **Error Type :**

   - In go, Error are values. The 'error' type is an interface that has a single method, 'Error()'

   ```go
    type error interface {
        Error() string
    }
   ```

2. **Creating Error :**

   - The error package provide a simple way to create a new error values

   ```go
    import "errors"

    var err := errors.New("an error occurred")
   ```

3. **Returning an Error :**

   - Functions that might encounter an error return an error value as the last return value.

   ```go
   func divide(a,b int)(int, error){
    if b==0{
        return 0, errors.New("Division by zero is not allowed")
    }
    return a / b, nil
   }
   ```

## Error handling patterns

1.  Checking for errors

    - After calling a function that returns an error u must check if the error is nil or not

    ```go
     a,b := 34,56
     res,err := divide(a,b)
     if err != nil {
         fmt.Println("an error occurred :",err)
     }else{
         fmt.Println("result : ",res)
     }

    ```

2.  Custom Error Types:

    - you can Define custom error types to provide more context

    ```go
     type myError struct {
         when time.time
         what string
     }

     func (e *myError) error() string{
         return fmt.Sprintf("at %v, %s",e.when,e.what)
     }

     func run() error{
         return &MyError{
             time.Now(),
             "it didnt work",
         }
     }

     func main(){
         if err := run(); err != nil{
             fmt.Println(err)
         }
     }
    ```

3.  Wrapping and Unwrapping Errors:

    - The fmt.Errorf function can wrap an existing error with additional context.

    ```go
    import "fmt"

    func readFile(fileName string) error {
    return fmt.Errorf("failed to read %s: %w", fileName, originalError)
    }
    ```

    To retrieve the original error, you can use the errors.Unwrap function.

    ```go

    import "errors"


    originalError := errors.New("original error")
    wrappedError := fmt.Errorf("wrapped error: %w", originalError)

    fmt.Println(errors.Unwrap(wrappedError)) // Output: original error
    ```

4.  Error Assertions:

    You can use type assertions to check for specific error types.

    ```go

    if err, ok := err.(\*MyError); ok {
    fmt.Println("Custom error occurred at", err.When)
    }
    ```

5.  Using errors.Is and errors.As:

        - errors.Is checks if an error is equal to a specific error.

    ```go
    Copy code
    if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File does not exist")
    }
    ```

    errors.As checks if an error can be cast to a specific type.

    ```go
    Copy code
    var pathError \*os.PathError
    if errors.As(err, &pathError) {
    fmt.Println("Failed at path:", pathError.Path)
    }
    ```
