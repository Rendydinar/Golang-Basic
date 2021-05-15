package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	AlamatRendy := Address{"Waingapu", "Nusa Tenggara Timur", "Indonesia"}
	// AlamatRendyClone := &AlamatRendy
	var AlamatRendyClone *Address = &AlamatRendy
	// Rendy pindah alamat 
	AlamatRendyClone.City = "Tokyo"

	// AlamatRendyClone = &Address{"Wangga", "Sumba Timur", "Indonesia"} // Ini membuat data baru tanpa mereferensikan ke alamat memori variabel yang diambil.

	*AlamatRendyClone = Address{"Wangga", "Sumba Timur", "Indonesia"} // Ini mmembuat data baru sekaligus memaksa address dari variabel yang ditampung untuk menggunakan data baru
	fmt.Println(AlamatRendy)
	fmt.Println(AlamatRendyClone)

	var AlamatSekolah *Address = new(Address) // membuat pointer dengan data yan kosong

	AlamatSekolah.City = "Jalan Supratman"
	fmt.Println(AlamatSekolah)

}

