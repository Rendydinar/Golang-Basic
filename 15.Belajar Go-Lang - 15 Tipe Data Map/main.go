package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{		
	// }
	person := map[string]string{
		"name": "Ninja Coder",
		"hobby": "Make Code",
	}

	person["title"] = "Engineering"
	fmt.Println("person", person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Golang For Ninja"
	book["author"] = "Master Golang"
	book["ups"] = "Salah"

	fmt.Println("book", book)
	fmt.Println("len book", len(book))
	delete(book, "ups")
	fmt.Println("book", book)
	fmt.Println("len book", len(book))
}