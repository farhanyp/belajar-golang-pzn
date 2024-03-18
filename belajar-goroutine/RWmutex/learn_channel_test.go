package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type accountBank struct{
	RWMutex sync.RWMutex
	Balance int
}

func (accountBank *accountBank) AddBalance(balance int){
	accountBank.RWMutex.Lock()
	accountBank.Balance += balance
	accountBank.RWMutex.Unlock()
}

func (accountBank *accountBank) GetBalance()int{
	accountBank.RWMutex.RLock()
	balance := accountBank.Balance
	accountBank.RWMutex.RUnlock()

	return balance
}

func TestGoroutine( t *testing.T){

	account := accountBank{}

	for i := 1; i <= 100; i++ {
		
		go func ()  {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}	
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.Balance)

}