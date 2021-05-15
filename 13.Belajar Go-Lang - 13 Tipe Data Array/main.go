package main

import "fmt"

func main() {
	var daftarNama [3]string
	daftarNama[0] = "Ninja Coder"
	daftarNama[1] = "wnnaCry"
	daftarNama[2] = "Poison"

	fmt.Println(daftarNama[0])
	fmt.Println(daftarNama[1])
	fmt.Println(daftarNama[2])

	// Membuat Array Secara Langsung
	var daftarNamaClone = [3]string{
		"Ninja Coder",
		"wnnaCry",
		"Poison",
	}
	fmt.Println(daftarNamaClone)

	var daftarNamaKosong [3]string

	fmt.Println("Jumlah Daftar Nama", len(daftarNama))
	fmt.Println("Jumlah Daftar Nama Clone", len(daftarNamaClone))
	fmt.Println("Jumlah Daftar Nama Kosong", len(daftarNamaKosong))

}