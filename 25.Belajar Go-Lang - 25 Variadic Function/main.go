package main

import "fmt"

func sumAll(numbers ...int) int{
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func main() {
	total := sumAll(50, 23, 23, 35, 123)
	fmt.Println(total)

	numbers := []int{50, 2332, 2233, 35, 123}
	total2 := sumAll(numbers...)
	fmt.Println(total2)

}