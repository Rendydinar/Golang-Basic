package helper

import "fmt"

// Bisa diakses di luar file
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// Tidak bisa diakses di luar file
func sayHello(name string) {
	fmt.Println("Hello", name)
}

const version = "1.0.0" // Tidak bisa diakses dari luar file
const Application = "Golang"  // Bisa diakses dari luar file