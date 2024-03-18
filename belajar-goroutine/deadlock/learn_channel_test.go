package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct{
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("User1 Has Lock")
	user1.Change(-amount)

	user2.Lock()
	fmt.Println("User2 Has Lock")
	user2.Change(amount)

	time.Sleep( 1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}


func TestGoroutine( t *testing.T){

	user1 := UserBalance{
		Name: "Maman",
		Balance: 100000,
	}

	user2 := UserBalance{
		Name: "Asep",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 10000)

	time.Sleep( 5 * time.Second)

	fmt.Println("User",user1.Name, ",", "Balance ", user1.Balance)
	fmt.Println("User",user2.Name, ",", "Balance ", user2.Balance)
}