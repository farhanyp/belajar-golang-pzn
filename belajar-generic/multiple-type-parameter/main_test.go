package multipletypeparameter

import (
	"fmt"
	"testing"
)

func Length[T any, K any](param1 T, param2 K) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestSample(t *testing.T){

	Length[int, string](100 ,"yp")

	Length[string, int]("yp", 100)

}