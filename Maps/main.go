package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["john"] = 30
	ages["alex"] = 45

	_,exists := ages["poly"]
	if exists {
		fmt.Println("age of poly is ",ages["poly"])
	}else{
		ages["poly"] = 45
	}

	for key, value := range ages {
		fmt.Println(key ," : ",value)
	}
}