# Defer, Panic dan Recover

## Defer
- Function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
- Akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi

## Panic
- Function yang digunakan untuk menghentikan program
- Biasa dipanggil saat menemukan error saat program berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

## Recover
- Function yang bisa digunakan untuk menanggkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.