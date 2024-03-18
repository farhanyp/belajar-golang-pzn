package main

import (
	"testing"
	"time"
)

func OnlyIn(channel chan <- string) {
	channel <- "Farhan Yudha Pratama"
}

func OnlyOut(channel <- chan string) {
	println(<- channel)
}

func TestGoroutine( t *testing.T){
	channel := make(chan string)

	go OnlyIn(channel)

	go OnlyOut(channel)

	time.Sleep(2 * time.Second)

	go close(channel)
}