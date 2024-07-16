package main

import "fmt"

type person struct {
	name string
	age  int
}

func(p *person) changeAge(age int){
	p.age = age
}

func main() {
	p := person{
		name: "xyz",
		age:  12,
	}
	fmt.Printf("Name : %v, Age : %v\n",p.name,p.age)
	fmt.Println("before changing age ",p.age)
	p.changeAge(34)
	fmt.Println("after changing age",p.age)
}