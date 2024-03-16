package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubTest( t *testing.T){

	t.Run("test1", func(t *testing.T){
		result := HelloWorld("yp")
		assert.Equal(t, "Hi yp", result)
	})

	t.Run("test2", func(t *testing.T){
		result := HelloWorld("yp")
		assert.Equal(t, "Hi yp", result)
	})

}