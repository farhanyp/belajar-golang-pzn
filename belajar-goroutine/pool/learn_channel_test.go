package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine( t *testing.T){
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("farhan")
	pool.Put("yudha")
	pool.Put("pratama")

	for i := 0; i < 10; i++{
		go func ()  {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)

}