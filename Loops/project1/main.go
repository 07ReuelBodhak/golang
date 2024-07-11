package main

import "fmt"

func main() {
	var start int
	fmt.Println("Enter the start number: ")
	fmt.Scan(&start)
	var end int
	fmt.Println("Enter the ending number: ")
	fmt.Scan(&end)

	fmt.Println("Prime numbers between", start, "and", end, "are:")
	for i := start; i <= end; i++ {
		if i < 2 {
			continue
		}
		isPrime := true
		for j := 2; j <= i/2; j++ {  
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
}
