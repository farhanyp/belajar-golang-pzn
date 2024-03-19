package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup){
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestGoroutine( t *testing.T){

	// Wait Group merupakan sebuah proses untuk mengunggu goroutine
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++{
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Complete")

}