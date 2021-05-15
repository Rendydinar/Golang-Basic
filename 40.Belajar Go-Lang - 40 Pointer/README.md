# Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

## Pass by value
- Secara default di Go-Lang semua variabel itu di passing by value, bukan by reference.

## Operator &
- Untuk membuat sebuah variabel dengan nilai pointer ke variabel yang lain, kita bisa menggunakan operator & diikuti dengan nama variabel nya.  

## Operator *
- Saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut.
- Semua variabel yang mencau ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut, bisa gunakan operator *

# Semoga tidak pusing ya :D