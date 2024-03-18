package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine( t *testing.T){
	channel := make(chan string, 3)

	go func(){
		channel <- "farhan"
		channel <- "yudha"
		channel <- "pratama"
	}()

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)

	time.Sleep(2 * time.Second)

	go close(channel)
}