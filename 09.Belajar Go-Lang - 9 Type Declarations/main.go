package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRendy NoKTP = "53496050100002"
	var statusRendy Married = false
	fmt.Println(noKtpRendy)
	fmt.Println("Menikah", statusRendy)
}