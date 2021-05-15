package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var enambulanpertama = months[0:6]

	fmt.Println("enambulanpertama", enambulanpertama)
	fmt.Println("len slice enambulanpertama", len(enambulanpertama))
	fmt.Println("kapasitas slice enambulanpertama", cap(enambulanpertama))

	// months[5] = "Juni_Dirubah"
	// fmt.Println(months)
	// enambulanpertama[4] = "Mei_Lagi"
	// fmt.Println(months)

	// var slice2 = months[2:4]
	var slice2 = months[11:]
	fmt.Println("slice2", slice2)

	var slice3 = append(slice2, "Ninja Coder")
	fmt.Println("slice3", slice3)	
	slice3[1] = "Bukan Desember"
	fmt.Println("slice3", slice3)	
	fmt.Println("slice2", slice2)
	fmt.Println("months", months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ninja"
	newSlice[1] = "Coder"
	fmt.Println("newSlice", newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice);
	fmt.Println("copySlice", copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("iniArray", iniArray)
	fmt.Println("iniSlice", iniSlice)
}  