package main

import "fmt"

type Man struct {
	Name string
}

// * = method struct pass by reference
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

	// fmt.Println(man) // hanya berlaku didalam function jika pass by value
}

func main() {
	rendy := Man{"Rendy"}
	rendy.Married()

	fmt.Println(rendy.Name)
}