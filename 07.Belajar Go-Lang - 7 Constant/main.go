package main

import "fmt"

func main() {
	// const hero string = "Hayabusa"
	// const Role = "Assassin"
	// const Gold = 300

	const (
		hero string = "Hayabusa"
		tipe = "Assassin"
		max_Health = 1000
	)

	fmt.Println(hero)
	fmt.Println(tipe)
	fmt.Println(max_Health)
}