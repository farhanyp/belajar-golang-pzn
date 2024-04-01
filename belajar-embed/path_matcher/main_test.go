package main

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed *.text
var file embed.FS

func TestEmbed(t *testing.T){

	dirs, _ := file.ReadDir(".")

	for _, dir := range dirs{
		if !dir.IsDir(){
			fmt.Println(dir.Name())
			content, _ := file.ReadFile(dir.Name())
			fmt.Println(string(content))
		}
	}

}
