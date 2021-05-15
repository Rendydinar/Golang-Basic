package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println(name, "Kamu di blok")
	} else {
		fmt.Println("Hallo", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Ninja Coder", blacklist)
	registerUser("Ninja", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}