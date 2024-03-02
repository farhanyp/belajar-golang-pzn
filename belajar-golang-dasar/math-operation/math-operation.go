package main

import "fmt";

func main(){

	// Math Operation
	var a int = 10
	var b int = 2
	var c int = 8
	var d int = 2
	var result int

	result = a + b - c / d * 20 % 2
	fmt.Println("math: ", result)

	// Unary Operator
	var i int = 1
	i++
	i++
	i--
	fmt.Println("Unary: ", i)

}