package main

import (
	"testing"
)


func BenchmarkHelloWord1(b *testing.B) {
	for i := 0 ; i < b.N; i++{
		HelloWorld("farhan")
	}
	
}

func BenchmarkHelloWord2(b *testing.B) {
	for i := 0 ; i < b.N; i++{
		HelloWorld("yp")
	}
	
}