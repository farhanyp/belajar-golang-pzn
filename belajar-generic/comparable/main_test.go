package comparable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](param1 T, param2 T) bool {

	if param1 == param2 {
		return true
	}else{
		return false
	}
}

func TestSample(t *testing.T){

	assert.True(t, IsSame[int](100 ,100))
	assert.True(t, IsSame[string]("yp" ,"yp"))

}