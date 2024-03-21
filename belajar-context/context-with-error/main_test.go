package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

)

func CreateCounter(ctx context.Context)chan int{
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select{
			case <-ctx.Done():
				// fmt.Println("berhenti")
				return
			default:
				// fmt.Println("lanjut")
				destination <- counter
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return destination
}

func TestContext(t *testing.T){
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine:", runtime.NumGoroutine())
}