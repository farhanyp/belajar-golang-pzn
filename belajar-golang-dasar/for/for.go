package main

import "fmt"

func main() {

	// var iteration int = 1
	// for iteration <= 3 {
	// 	fmt.Println("ini", iteration)
	// 	iteration ++
	// }

	// for dengan statement
	// for counter := 1; counter <=3; counter++{
	// 	fmt.Println("ini", counter)
	// }

	// for untuk collection
	// names := []string{"farhan", "yudha", "pratama"}
	// for index, name := range names{
	// 	fmt.Println(index, ": ", name)
	// }

	// for untuk collection tanpa index
	names := []string{"farhan", "yudha", "pratama"}
	for _ , name := range names{
		fmt.Println(name)
	}
}