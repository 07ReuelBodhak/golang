package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
}

type Circle struct {
	radius int
}

type Rectangle struct {
	l int
	b int
}

func (c Circle) area() float32 {
	pi := math.Pi
	return float32(pi) * float32(c.radius * c.radius)
}

func (r Rectangle) area() float32{
	return float32(r.l * r.b)
}

func describeValue(i interface{}){
	fmt.Printf("type %T : value %v\n",i,i)
}

func main() {
	r := Rectangle{
		l: 34,
		b: 54,
	}
	fmt.Println("area of rectangle ",r.area())
	c := Circle{
		radius: 43,
	}
	fmt.Println("area of circle ",c.area())

	mystery := interface{}(rune(32))
	describeValue(mystery)

	s,ok := mystery.(string)
	if ok {
		fmt.Println(s) 
	} else {
		fmt.Println("mystery is not a string")
	}

	switch v := mystery.(type) {
	case string:
		fmt.Println("mystery is a string:", v)
	case int:
		fmt.Println("mystery is an int:", v)
	default:
		fmt.Printf("mystery is of type %T\n", v)
	}
}
