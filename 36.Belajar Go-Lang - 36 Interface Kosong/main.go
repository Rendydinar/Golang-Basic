package main

import "fmt"

func Ups(i int, apapun interface{}) interface{} {
	if i == 1 {
		return 2
	} else if i == 2 {
		return true
	} else {
		return "Upsh"
	}
}

func main() {
	var data  interface{}= Ups(3)
	fmt.Println(data)
}