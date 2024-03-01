package main

import "fmt";

func main(){

	// Membuat variabel mengunakan type datas
	var variabel1 string = "Farhan Yudha Pratama"
	fmt.Println(variabel1)

	// Membuat variabel dengan mengunakan type data otomatis
	var variabel2 = "Farhan Yudha Pratama"
	fmt.Println(variabel2)

	// Membuat variabel dengan simple
	// Tetapi metode ini hanya untuk deklarasi pertama
	variabel3 := "Farhan Yudha Pratama"
	fmt.Println(variabel3)

	// Membuat banyak variabel sekaligus
	var (
		variabel4 string = "Farhan Yudha Pratama"
		variabel5 string = "Farhan Yudha Pratama"
	)
	fmt.Println(variabel4)
	fmt.Println(variabel5)
}