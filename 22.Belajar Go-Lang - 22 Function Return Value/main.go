package main

import "fmt"

// function return data string
func getHello(name string) string {
	if name == "" {
		return "Hello bro, Siapa Namamu ?"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Ninja Coder")
	fmt.Println(result)
}