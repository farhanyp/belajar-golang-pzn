package main

import "fmt"

func main() {

	// Break berfungsi untuk menghentikan perulangan
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan ke ", i)
	}

	// Continue berfungsi untuk menjeda perulangan
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke ", i)
	}
}