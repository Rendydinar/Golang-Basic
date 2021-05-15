# Package Initialization

- Saat kita membuat oackage, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
- Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database
- Untu membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

# Blank Identifier

- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
- Secara default, Go-Lang akan komplen ketika pakcage yang diimport namun tidak digunakan 
- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier(_) sebelum nama package ketika melakukan import 