package main

import "fmt"

// Return multiple value tipe string
func getFullName()(string, string) {
	return "Ninja", "Coder"
}

func main() {
	// firstName ambil return pertama, lastName ambil return kedua
	// firstName, lastName := getFullName();

	// ignore data return kedua
	firstName, _ := getFullName(); 

	fmt.Println(firstName)
}