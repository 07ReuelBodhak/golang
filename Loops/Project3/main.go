package main

import (
	"fmt"
)

func main() {
	var first int
	fmt.Println("enter first number:")
	fmt.Scan(&first)

	var second int
	fmt.Println("enter second number")
	fmt.Scan(&second)

	for {
		max_num := max(first,second)
		min_num := min(first,second)

		c := max_num % min_num
		if c == 0 {
			fmt.Printf("greatest common divisor is: %d",min_num)
			break
		}
		first = min_num
		second = c
	}

}

