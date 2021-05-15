package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func main() {
	var NinjaCoder Customer
	NinjaCoder.Name = "Hayabusa"
	NinjaCoder.Address = "Japan"
	NinjaCoder.Age = 25

	fmt.Println(NinjaCoder.Name)
	fmt.Println(NinjaCoder.Address)
	fmt.Println(NinjaCoder.Age)

	fmt.Println(NinjaCoder)

	jokowi := Customer{
		Name: "Jokowi",
		Address: "Surakarta",
		Age: 59,
	}
	fmt.Println(jokowi)

	viktorLaiskodat := Customer{"Viktor Laiskodat", "Oenesu", 56}
	fmt.Println(viktorLaiskodat)

}