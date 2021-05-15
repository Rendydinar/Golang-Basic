package main

import (
	"fmt"
	_"example.com/database"
) // package database diakses secara blank identifier
func main() {
	result := database.GetDatabase();

	fmt.Println(result)
}