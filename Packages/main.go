package main

import (
	"fmt"
	"packages/calculator"
)

func main() {
	sum := calculator.Add(2,3)
	difference := calculator.Sub(4,6)
	fmt.Println("sum is ",sum)
	fmt.Println("difference is",difference)
}