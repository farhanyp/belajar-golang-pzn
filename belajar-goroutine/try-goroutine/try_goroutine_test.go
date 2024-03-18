package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestGoroutine( t *testing.T){

	go HelloWorld()
	fmt.Println("goroutine")

	time.Sleep( 1 * time.Second)
}