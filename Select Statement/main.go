package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	quit := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "operation complete"
	}()
	go func(){
		time.Sleep(3 * time.Second)
		quit <- true
	}()
	for {
		select{
		case res := <- ch:
			fmt.Println(res)
		case <- quit:
			fmt.Println("time out")
			return
		}
	}
	
}