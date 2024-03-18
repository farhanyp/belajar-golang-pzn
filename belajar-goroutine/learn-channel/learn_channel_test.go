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
	channel := make(chan string)

	go func ()  {
		time.Sleep( 2 * time.Second)
		channel <- "Farhan Yudha Pratama"
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep( 5 * time.Second)
}