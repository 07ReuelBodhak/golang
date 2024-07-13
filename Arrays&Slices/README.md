# Arrays and Slices

# Arrays

- An array is a collection of element of same datatype , with fixed sized defined at the time of declaration

- syntax :
  ```go
  var arrayName [size]Type
  ```

## Accessing Array element

- An array element are accessed using the index, which starts at 0

- Example :

  ```go
  numbers[0] = 10
  fmt.Println(numbers[0])
  ```

## Iterating over array

- You can use a _for_ loop to iterate over an array

- Example :

  ```go
  for i :=0 ; i< len(numbers); i++{
      fmt.Println(numbers[i])
  }
  ```

## Array Characteristic

- Array have a fixed sized
- the size of the array is part of its type
- Arrays are value types. Assigning one array to another copies all elements

# Slices in go

- A slice is dynamically-sized, flexible view into the elements of an array

- Syntax :

  ```go
  var sliceName []Type
  ```

## Create a slice

- we can create slice using _make_ built in function

  ```go
    numbers := make([]int,5) // creates a slice with length 5 and capacity 5
  ```

## appending elements in a slice

    lets say we have a slice of numbers

    ```go
    numbers := []int{100,23,56}
    ```

    no we want to append 67 in a numbers

    ```go
    append(numbers,67)
    ```
    now this will not modify the original slice this will create a new numbers slice so to modify original slice we will type like this

    ```go
    numbers = numbers.append(numbers,67)
    ```

## iterating over slices

we can iterate over an array using range function or using for loop

- using for loop

  ```go
  for i:=0 ; i < len(numbers); i++{
    fmt.Println(i)
  }
  ```

- using range function

  ```go
  for i,num := range num{
    fmt.Println(i,num)
  }
  ```
