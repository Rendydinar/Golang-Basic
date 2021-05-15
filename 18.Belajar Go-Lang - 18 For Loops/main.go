package main

import "fmt"

func main() {
	// counter := 1
	// for counter := 1 {
	// 	fmt.Println("Perulangan ke -", counter)
	// 	counter++
	// }

	// For dengan statement
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke -", counter)
	// }

	names := []string{"wnnaCry", "Ninja Coder", "r3ndy"}
	for counter := 0; counter < len(names); counter++ {
		fmt.Println("names[",counter,"] = ", names[counter])
	}

	for _, name := range names {
		// fmt.Println("Index", i, "=", name)
		fmt.Println(name)
	}

	hero := make(map[string]string)
	hero["name"] = "Hayabusa"
	hero["type"] = "Assasin"

	for key, value := range hero {
		fmt.Println(key, "=", value)
	}
}