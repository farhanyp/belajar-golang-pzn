package typeparameter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)

	return param
}

func TestSample(t *testing.T){

	var resultInt = Length[int](100)
	assert.Equal(t, 100, resultInt)

	var resultString = Length[string]("yp")
	assert.Equal(t, "yp", resultString)

}