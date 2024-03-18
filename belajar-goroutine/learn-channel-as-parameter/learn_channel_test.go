package main

import (
	"testing"
	"time"
)

func HelloWorld(channel chan string) {
	channel <- "Farhan Yudha Pratama"
}

func TestGoroutine( t *testing.T){
	channel := make(chan string)

	go HelloWorld(channel)

	println(<- channel)

	time.Sleep(1 * time.Second)

}