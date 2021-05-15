package main

import "fmt"

func main() {
	var namaAsli = "Ninja Coder"
	var namaPalsu = "Palsu Ninja Coder"

	fmt.Println(namaAsli == namaPalsu)
	fmt.Println(namaAsli < namaPalsu)

	var value1 = 100
	var value2 = 100

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}