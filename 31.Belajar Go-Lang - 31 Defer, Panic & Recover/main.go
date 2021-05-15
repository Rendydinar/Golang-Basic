package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 18 / value
	fmt.Println("Result", result)
}

func endApp() {
	// Testing recover
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func runAppWithRecover(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	// Testing defer
	// runApplication(0)
	// Testing panic
	// runApp(true)
	// Testing recover
	runAppWithRecover(true)
	fmt.Println("Ninja Coder")
}