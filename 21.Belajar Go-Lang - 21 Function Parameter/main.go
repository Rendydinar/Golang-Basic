package main

import "fmt"

func sayHelloTo(name string, info string) {
	fmt.Println("Hello", name, "you are", info)
}

func main() {
	sayHelloTo("Ninja Coder", "Awesome")
}