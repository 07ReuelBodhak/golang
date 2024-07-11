package main

import "fmt"

func main() {
	var number int
	fmt.Println("enter number to print pattern: ")
	fmt.Scan(&number)
	fmt.Println()

	for i := 1; i <= number; i ++ {
		for j:= 0; j < i; j ++ {
			fmt.Printf("%d",i)
		}
		fmt.Println()
	}
}