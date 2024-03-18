package main

import (
	"fmt"
	"testing"
)

func GiveMeChannel1 (channel1 chan string){
	channel1 <- "Channel 1"
}

func GiveMeChannel2 (channel2 chan string){
	channel2 <- "Channel 2"
}

func TestGoroutine( t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeChannel1(channel1)
	go GiveMeChannel2(channel2)

	counter := 0
	for{
		select{
		case data := <- channel1:
			fmt.Println("Data dari channel-1 : ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel-2 : ", data)
			counter++
		}

		if counter == 2 {
			break
			
		}
	}

}