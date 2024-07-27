package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(email int,ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for i :=0; i < email; i ++{
		fmt.Println("sending message ",i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func receiveEmail(ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for i := range ch{
		fmt.Println("email received", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	ch := make(chan int)
	wsg := &sync.WaitGroup{}
	wsg.Add(2)
	go sendEmail(12,ch,wsg)
	go receiveEmail(ch,wsg)
	wsg.Wait()
}