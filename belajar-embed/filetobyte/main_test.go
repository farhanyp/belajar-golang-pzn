package main

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed halal.jpeg
var file []byte

func TestEmbed(t *testing.T){

	err := ioutil.WriteFile("halal_next", file, fs.ModePerm)
	if err != nil {
		panic(err)
	}

}

