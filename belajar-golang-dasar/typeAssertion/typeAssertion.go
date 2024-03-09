package main

import (
	"fmt"
)

// contoh function yang mengembalikan any
func random() any {
	return 1
}

func main() {

	random := random()

	// check type data
	switch random.(type){
	case string:
		fmt.Println("ini String")
	case int:
		fmt.Println("ini Int")
	default:
		fmt.Println("ini Unkown")
	}


}