package main

import "fmt"

func rekursive(number int) int{
	if number == 1 {
		return 1
	} else {
		return number * rekursive(number-1)
	}
}

func main() {
	fmt.Println(rekursive(5))
}