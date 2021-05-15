# Tipe Data Map
- Adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang kita gunakan
- Sederhananya, Map adalah tipe data kumpulan key-value(kata kunci-nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukan ke dalam Map boleh sebanyak-banyakny, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

## Function Map
- len(map) - Mendapatkan jumlah data di map
- map[key] - Mengambil data di map dengan key
- map[key] = value - Mengubah data di map dengan key
- make(map[TypeKey]TypeValue) - Membuat map baru
- delete(map, key) - Menghapus data di map dengan key