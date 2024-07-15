package main

import "fmt"

type player struct {
	health int
	element int
}

func takeDamageFromExplosion(p *player) {
	fmt.Println("player is taking damage from explosion")
	explosionDamage := 10
	if p.health < 10 {
		fmt.Println("player died")
	}
	p.health -= explosionDamage
}

func (p *player) addElement(elem int){
	fmt.Println("new element added")
	p.element += elem
}

func main() {
	player := player{
		health: 100,
		element:  0,
	}
	fmt.Println("health before explosion : ",player.health)
	takeDamageFromExplosion(&player)
	fmt.Println("health after explosion : ",player.health)
	fmt.Println("before adding element", player.element)
	player.addElement(34)
	fmt.Println("after adding element", player.element)
}