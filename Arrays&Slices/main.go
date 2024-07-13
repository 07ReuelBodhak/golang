package main

import (
	"fmt"
)

func main() {

	//declaring array
	var ages = [3]int{12, 34, 66}
	fmt.Println(ages)

	names := [2]string{"john","alex"}
	fmt.Println(names)

	percentage := [4]float32{43.56,67.89}
	percentage[2] = 78.45
	percentage[3] = 98.33
	fmt.Println(percentage)

	//slices 
	var scores = []int{100,50,60}
	
	scores = append(scores, 65)
	fmt.Println(scores)

	for i,num := range scores{
		fmt.Printf("%d : %d,",i,num)
	}
	
	//using slice operator
	rangeOne := scores[:2]
	fmt.Println(rangeOne)

	rangeTwo := percentage[1:3]
	fmt.Println(rangeTwo)
}