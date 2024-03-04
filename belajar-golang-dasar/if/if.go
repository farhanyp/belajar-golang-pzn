package main

import "fmt"

func main() {

	name:= "rahmat"

	if name == "farhan"{
		fmt.Println("ini adalah farhan")
	}else if name == "rahmat"{
		fmt.Println("ini adalah rahmat")
	}else{
		fmt.Println("bukan farhan dan rahmat")
	}

	// if short statement
	if length := len(name); length == 6 {
		fmt.Println("ini adalah farhan karena memiliki 6 char")
	}else{
		fmt.Println("ini bukan farhan")
	}

}