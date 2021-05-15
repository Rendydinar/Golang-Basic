# Tipe Data Slice
- Potongan dari data Array
- Mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

## Detail Tipe Data Slice
- Tipe Data Slice memiliki 3 data, pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array pada slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

## Membuat Slice Dari Array
- array[low:high] - Membuat slice dari array dari, dimulai index low sampai index sebelum high
- array[low:] - Membuat slice dari array dimulai dari index low sampai index akhir di array
- array[:high] - Membuat slice dari array dimulia index 0 sampai index sebelum high
- array[:] - Membuat slice dari array dimulai index 0 sampai index di akhir array

## Function Slice
- len(slice) - Mendapatkan panjang
- cap(slice) - Mendapatkan kapasitas
- append(slice, data) - Membuat slice baru dengan menambah data ke posisi terkahir slice, jika kapasitas sudah penuh, maka akan membuat array baru
- make([]TypeData, length, capacity) - Membuat slice baru
- copy(destination, source) - Menyalin slice dari source ke destination