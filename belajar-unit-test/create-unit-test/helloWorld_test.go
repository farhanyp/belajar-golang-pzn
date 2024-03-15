package main

import "testing"

func TestHelloWorld( t *testing.T){

	result := HelloWorld("yp")

	if result != ("Hello yp"){
		panic("Result not match")
	}

}