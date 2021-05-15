package database

var connection string

// Deklarasi function Initialization
func init () {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}