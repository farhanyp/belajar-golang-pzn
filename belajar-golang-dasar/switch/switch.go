package main

import "fmt"

func main() {

	name:= "farhan"

	switch name {
	case "farhan":
		fmt.Println("ini adalah farhan")
	case "rahmat":
		fmt.Println("ini adalah farhan")
	default:
		fmt.Println("ini bukan farhan dan rahmat")
	}

	// switch short statement
	switch length := len(name); length == 6 {
	case true:
		fmt.Println("ini adalah farhan karena memiliki 6 char")
	case false:
		fmt.Println("ini bukan farhan")
	}


	// switch tanpa kondisi
	length := len(name)
	switch{
	case length == 6:
		fmt.Println("ini adalah farhan karena memiliki 6 char")
	case length < 6:
		fmt.Println("ini bukan farhan")
	}

}