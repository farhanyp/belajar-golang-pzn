package main

import (
	"testing"
)


func BenchmarkHelloWord1(b *testing.B) {
	b.Run("test1", func(b *testing.B){
		for i := 0 ; i < b.N; i++{
			HelloWorld("test1")
		}
	})

	b.Run("test2", func(b *testing.B){
		for i := 0 ; i < b.N; i++{
			HelloWorld("test2")
		}
	})
	
}