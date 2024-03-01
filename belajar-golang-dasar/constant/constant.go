package main

import "fmt";

func main(){

	// Membuat contant mengunakan type data
	const contant1 string = "Farhan Yudha Pratama"
	fmt.Println(contant1)

	// Membuat contant mengunakan type data
	const contant2 string = "Farhan Yudha Pratama"
	fmt.Println(contant2)

	// Membuat banyak constant sekaligus
	const (
		constant3 string = "Farhan Yudha Pratama"
		constant4 string = "Farhan Yudha Pratama"
	)
	fmt.Println(constant3)
	fmt.Println(constant4)
}