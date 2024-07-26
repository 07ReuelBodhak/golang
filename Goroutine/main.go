package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for job := range ch {
		fmt.Printf("Worker %d started job %d\n",id,job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n",id,job)
	}
}

func main() {
	ch := make(chan int,10)

	for i:=0; i<3; i++{
		go worker(i, ch)
	}

	for j:=0; j<9; j++{
		ch <- j
		fmt.Printf("Job %d sent\n", j)
	}
	close(ch)

	time.Sleep(time.Second * 3)
}