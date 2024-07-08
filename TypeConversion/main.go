package main

import (
	"fmt"
	"strconv"
)

func main() {
	exp := 344
	var floatVal float32 = float32(exp)
	var byteVal byte = byte(exp)
	fmt.Printf("converting integer %v to float %.2f and rune %v\n",exp,floatVal,byteVal)

	exp2 := "123"
	integerVal,err := strconv.Atoi(exp2) 
	if err != nil {
		fmt.Println("cannot convert string to integer")
	}
	fmt.Printf("Converting string %v to integer %d",exp2,integerVal)
}
