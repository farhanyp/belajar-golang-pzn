package main

import (
	"testing"
)

func TestHelloWorldTable( b *testing.B){
	benchmarks := []struct{
		name string
		expected string
	}{
		{
			name: "yp1",
			expected: "Hello yp1",
		},
		{
			name: "yp2",
			expected: "Hello yp2",
		},
	}

	for _ , benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0 ; i < b.N; i++{
				HelloWorld(benchmark.expected)
			}
		})
	}

}