package main

import "fmt"

func main() {
	//short variable declaration
	h := 23.45
	fmt.Printf("short variable declaration %v\n",h)

	// Declaring multiple variable 
	var b,c int8
	b = -23
	c = -34
	fmt.Printf("Declaring multiple variable %v,%v\n",b,c)

	//Group Declaration
	var(
		e float32
		f string
	)
	e = 12.345
	f = "hello"
	fmt.Printf("Declaring multiple variable using var keyword %v,%v\n",e,f)

	//Declaring variable with initializer
	var g = 233
	fmt.Printf("Declaring variable with initializer %v\n",g)

	//Constant variable
	const i uint16 = 1235
		fmt.Printf("constant variable declaration %v",i)
}