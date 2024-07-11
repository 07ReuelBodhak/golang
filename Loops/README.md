# Loops in Go lang

in go lang there are three types of loops available :

1. for loop
2. while loop
3. infinite loop

### For loop

syntax :

```go
for initialization; condition; post {

}
```

- initialization means the from where the loop starts its execution
- condition means until where it will execute
- Post statement Executed at the end of each iteration. Often used to increment or update variables.

### While loop

syntax :

```go
for condition {

}
```

in this the for loop will execute until condition becomes false

### Infinite loop

syntax :

```go
for {

}
```

runs loop infinitely

### Project

1. **Prime Number Checker with Range**: Write a program that prompts the user to enter two integers (say start and end). The program should then print all prime numbers within that range. Use nested loops and a prime-checking algorithm to solve this problem efficiently. [link](./project1/main.go)

2. **Pattern Generation**: Create a program that generates a specific pattern based on user input (e.g., number of rows). For example, if the user enters 5, the program should generate a pattern like below using loops:

```
1
22
333
4444
55555
```

[link](./project2/main.go)

3. Matrix Operations: Implement a program that performs matrix multiplication. Prompt the user to enter the dimensions of two matrices (e.g., 2x3 and 3x2), and then input the elements for each matrix. Use loops to perform the multiplication according to matrix multiplication rules (sum of products of rows and columns).
