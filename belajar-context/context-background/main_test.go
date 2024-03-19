package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T){

	// mendefiniskan context kosong
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}