# Type Assertions 
- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan 
- Fitur ini sering sekali digunakan ketika bertemua dengan data interface kosong

## Type Assetions Menggunakan Switch
- Saat menggunakan type assertions, maka bisa berakibatkan terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan swithc expression untuk melakukan type assertions 