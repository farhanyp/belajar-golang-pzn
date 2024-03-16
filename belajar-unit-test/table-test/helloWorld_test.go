package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldTable( t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "yp1",
			request: "yp1",
			expected: "Hello yp1",
		},
		{
			name: "yp2",
			request: "yp2",
			expected: "Hello yp2",
		},
	}

	for _ , test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result  := HelloWorld(test.request)

			assert.Equal(t, test.expected, result)
		})
	}

}