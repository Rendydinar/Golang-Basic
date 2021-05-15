package main

import "fmt"

func main() {
	name := "wnnaCry"

	switch name {
	case "Ninja Coder": 
		fmt.Println("Hello Ninja Coder")
	case "wnnaCry": 
		fmt.Println("Hello wnnaCry")
	default: 
		fmt.Println("Nama Tidak Dikenal, Banned", name)
	}

	// switch lenght := len(name); lenght > 5 { // Short Statetment
	// case true:
	// 	fmt.Println("Nama yang cukup panjang")	
	// case false:
	// 	fmt.Println("Nama yang bagus")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")	
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama yang bagus")		
	}
}