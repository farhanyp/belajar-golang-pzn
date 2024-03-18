package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine( t *testing.T){

	x := 0
	var mutex sync.Mutex
	// Contoh race condition karena ada yang bertabarakan jadinya tidak sampai 100.000 perulangan
	// dengan menggunakan mutex bisa menghindari race condition
	for i := 1; i <= 1000; i++{
		go func ()  {
			for j := 1; j <= 100; j++{
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)

}