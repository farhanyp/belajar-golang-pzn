package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssertion( t *testing.T){
	result := HelloWorld("yp")

	assert.Equal(t, "Hi yp", result, "Actual output is Hi yp")

}