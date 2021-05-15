package main

import "fmt"

func getFullName()(firsname, middlename, lastname string) {
	firsname = "Umbu"
	middlename = "Theofilus"
	lastname = "Dendimara"

	return
}

func main() {
	a, b, c := getFullName();
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}