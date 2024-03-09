package database

var connection string

// Init akan dijalankan ketika package dipanggil
func init() {
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}