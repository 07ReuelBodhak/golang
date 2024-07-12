package main

import "fmt"

type Rectangle struct{
	width int
	height int
}

//basic function

func Add(x,y int) int{
	return x + y
}

//Multiple Return value
func divide(x,y int) (int,int){
	qutioent := x / y
	reminder := x % y
	return qutioent, reminder
}

//named Return value
func somestuff(sum int)(x,y int){
	x = x + sum 
	y = y * sum
	return
}

//Variadic function

func Total_Sum(nums ...int) int{
	total := 0
	for _,num := range nums{
		total += num
	}
	return total
}

//High order function
func Apply(f func(int) int, value int) int{
	return f(value)
}

func Square(n int) int{
	return n * n
}


//methods

func (r Rectangle) Area()int{
	return r.height * r.width
}

func main(){
	sum := Add(2,4)
	fmt.Println("result of addition : ",sum)

	que,rem := divide(12,5)
	fmt.Printf("Qutioent %v Reminder %v \n",que,rem)

	x,y := somestuff(233)
	fmt.Printf("after something to x %d, after something to y %d\n",x,y)

	nums := []int{1,2,3,5,6,7}
	total := Total_Sum(nums...)
	fmt.Println("total are : ",total)

	//annonymus function 
	difference := func(a int,b int) int{
		return a - b
	}
	difer := difference(4,5)
	fmt.Println("difference of two number are",difer)

	square := Apply(Square,3)
	fmt.Println("square is",square)

	rect := Rectangle{
		height: 24,
		width: 45,
	}
	area := rect.Area()
	fmt.Println("Area of rectangle", area)

	//function literals
	result := func(x,y int) int{
		return x*y
	}(12,4)
	fmt.Println("Result of multiplication",result)
}