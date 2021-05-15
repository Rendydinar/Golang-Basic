package main

import "fmt"

func main() {
	var name string

	name = "Ninja Coder"
	fmt.Println(name)

	name = "Golang Ninja"
	fmt.Println(name)

	// tanpa deklarasi tipe data
	var friendName = "Hatori"
	fmt.Println(friendName)

	var age int8 =  20
	fmt.Println(age)

	// :=
	country := "Indonesia"
	fmt.Println(country)

	var (
		hero = "Hayabusa"
		tipe = "Assassin" 
	)
	fmt.Println(hero)
	fmt.Println(tipe)

}