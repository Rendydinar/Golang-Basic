# Error Interface
- Golang memiliki sebuah interface yang digunakan sebagai kontrak untuk membaut error, namanya interface nya adalah error
- Syntaks: 
	type error interface {
		Error() string
	}

## Membuat Error
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors.