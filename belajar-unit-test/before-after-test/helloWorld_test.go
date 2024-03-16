package main

import (
	"fmt"
	"testing"
)

func Testmain(m *testing.M){

	fmt.Println("Sebelum melakukan test")

	m.Run()

	fmt.Println("Sesudah melakukan test")

}

func TestHelloWorldyp1( t *testing.T){

	fmt.Println("Sedang melakukan test")

}