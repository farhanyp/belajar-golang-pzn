package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed file.text
var file string

func TestEmbed(t *testing.T){

	fmt.Println(file)

}

