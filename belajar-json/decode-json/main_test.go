package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct{
	Firstname, Middlename, Lastname string
	Age int
	Ismarried bool
}

func TestJson(t *testing.T) {

	jsonString := `{"Firstname":"Farhan","Middlename":"Yudha","Lastname":"Pratama","Age":23,"Ismarried":true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Firstname)
	fmt.Println(customer.Age)
	fmt.Println(customer.Ismarried)

}