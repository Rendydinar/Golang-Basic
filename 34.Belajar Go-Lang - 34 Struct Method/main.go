package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}


func main() {
	var NinjaCoder Customer
	NinjaCoder.Name = "Hayabusa"
	NinjaCoder.Address = "Japan"
	NinjaCoder.Age = 25

	fmt.Println(NinjaCoder.Name)
	fmt.Println(NinjaCoder.Address)
	fmt.Println(NinjaCoder.Age)

	NinjaCoder.sayHello("Hanabi")

}

