package config

import "fmt"

// Init akan dijalankan ketika package dipanggil
func init() {
	fmt.Println("This is Internal")
}