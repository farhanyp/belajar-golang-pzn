package main

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed file1.text
//go:embed file2.text
//go:embed file3.text
var file embed.FS

func TestEmbed(t *testing.T){

	a, _ := file.ReadFile("file1.text")
	fmt.Println(string(a))

	b, _ := file.ReadFile("file2.text")
	fmt.Println(string(b))

	c, _ := file.ReadFile("file3.text")
	fmt.Println(string(c))

}

