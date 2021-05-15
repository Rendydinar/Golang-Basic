package main

import "fmt"

func main() {
	counter := 0
	name := "Ninja Coder"

	fmt.Println("Sebelum akses fitur Closure")
	fmt.Println(counter)
	fmt.Println(name)
	increment := func() {
		fmt.Println("Increment")
		// Closure
		counter++
		name = "Ninja Palsu"
	}

	increment()
	increment()

	fmt.Println("Sesudah akses fitur Closure")
	fmt.Println(counter)
	fmt.Println(name)
}