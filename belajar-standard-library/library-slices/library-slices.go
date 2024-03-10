package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"Farhan", "Yudha", "Pratama", "Yp", "Maman"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Yudha"))
	fmt.Println(slices.Index(names, "Farhan"))
	fmt.Println(slices.Index(names, "Paul"))

}