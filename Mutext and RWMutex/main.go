package main

import (
	"fmt"
	"sync"
	"time"
)

type player struct {
	health 	int
	m      *sync.Mutex
}

func newPlayer() player{
	return player{
		health: 100,
		m : &sync.Mutex{},
	}
}

func (p *player)fire(wg *sync.WaitGroup){
	p.m.Lock()

	if p.health <= 0{
		fmt.Print("you died")
		p.m.Unlock()
		wg.Done()
		return
	}

	p.health -= 10
	fmt.Println("fire damage was taken your health : ",p.health)
	p.m.Unlock()
	time.Sleep(time.Millisecond * 400)
	wg.Done()
}

func (p *player) potion(wg *sync.WaitGroup){
	p.m.Lock()
	if p.health >= 100{
		fmt.Println("health is full")
		p.m.Unlock()
		wg.Done()
		return
	}

	p.health += 5
	fmt.Println("health after potion : ",p.health)
	p.m.Unlock()
	time.Sleep(time.Millisecond * 200)
	wg.Done()
}

func main() {
	p :=  newPlayer()
	var wg sync.WaitGroup

	for i:=0;i<10;i++{
		wg.Add(2)
		go p.fire(&wg)
		go p.potion(&wg)
	}
	wg.Wait()
}