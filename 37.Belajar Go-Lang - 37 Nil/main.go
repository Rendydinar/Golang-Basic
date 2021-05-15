package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	var person = NewMap("Ninja Coder")

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}
}