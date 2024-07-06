package main

import "fmt"

func main() {
	// signed Integer
	var a int8 = -128

	// unsigned Integer
	var b uint8 = 128

	// Alias types
	var c byte = 255
	var d rune = 'a'

	//floating point number
	var f float32 = 1.24

	//complex number
	var g complex64 = 1 + 2i
	var h complex128 = 2 - 3i

	fmt.Printf("Signed-Integer : %v \nUnsigned-Integer : %v \nByte : %v \nRune : %v\nFloating Number : %v\nComplex number 1 : %v\nComplex number 2 : %v",a,b,c,d,f,g,h)

}