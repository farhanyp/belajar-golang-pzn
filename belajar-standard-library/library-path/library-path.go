package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	fmt.Println(filepath.Dir("../library-os"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))

}