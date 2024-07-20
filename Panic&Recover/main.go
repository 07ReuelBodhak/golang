package main

import "fmt"

func printRes(res []string, children int) {
	defer func(){
		if r := recover(); r!= nil{
				fmt.Println("Recovered from panic : ",r)
		}
	}()

	for i := 0; i < children; i++ {
		if i == len(res){
			panic("not enough resources")
		}
		fmt.Printf("%d : %s\n",i+1,res[i])
	}
}

func main() {
	resources := []string{"apple", "bat", "ball", "kite"}
	children := 5
	printRes(resources,children)
}

