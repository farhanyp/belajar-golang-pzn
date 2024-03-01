package main

import "fmt";

func main(){

	// Konversi type data
	var nilai32 int32 = 1234
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	// Konversi type data
	var name string = "Farhan Yudha Pratama"
	var e uint8  = name[0]
	var eString string = string(e)
	fmt.Println(eString)
}