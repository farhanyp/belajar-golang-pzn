package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine( t *testing.T){

	x := 0

	// Contoh race condition karena ada yang bertabarakan jadinya tidak sampai 100.000 perulangan
	for i := 1; i <= 1000; i++{
		go func ()  {
			for j := 1; j <= 100; j++{
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)

}