package main

import "fmt"

func main() {

	// membuat array
	var array1 [3]string

	array1[0] = "farhan"
	array1[1] = "yudha"
	array1[2] = "pratama"

	fmt.Println(array1)

	// membuat array secara langsung
	var array2 = [3]string{
		"farhan",
		"yudha",
		"pratama",
	}
	fmt.Println(array2)

	// membuat array tanpa batas
	var array3 = [...]string{
		"farhan",
		"yudha",
		"pratama",
	}
	fmt.Println(array3)

	// function pada array
	array3[0] = "racing"
	fmt.Println(len(array3))
	fmt.Println(array3[0])
}