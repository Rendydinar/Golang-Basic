package main

import "fmt"

func main() {
	// var name = "wnnaCry"
	if name := "Ninja Coder"; name == "Ninja Coder" { // If dengan short statement
		fmt.Println("Nama Benar, Hallo Ninja Coder")
	} else if name == "wnnaCry" {
		fmt.Println("Nama Benar, Hallo wnnaCry")
	} else {
		fmt.Println("Nama Tidak Benar, Report", name)		
	}
}