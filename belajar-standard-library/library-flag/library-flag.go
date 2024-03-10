package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "localhost", "Enter your username")

	flag.Parse()

	fmt.Println(*username)
}