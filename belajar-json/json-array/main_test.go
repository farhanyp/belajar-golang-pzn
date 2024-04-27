package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct{
	Street string
}

type Customer struct{
	Firstname, Middlename, Lastname string
	Age int
	Ismarried bool
	Hobbies []string
	Address []Address
}

func TestEncodingJsonArray(t *testing.T) {

	customer := Customer{
		Firstname: "Farhan",
		Middlename: "Yudha",
		Lastname: "Pratama",
		Age: 23,
		Ismarried: true,
		Hobbies: []string{"Gaming", "Reading", "Learning"},
		Address: []Address{
			{
				Street: "jalan1",
			},
			{
				Street: "jalan2",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestDicodingJsonArray(t *testing.T) {

	jsonString := `{"Firstname":"Farhan","Middlename":"Yudha","Lastname":"Pratama","Age":23,"Ismarried":true,"Hobbies":["Gaming","Reading","Learning"],"Address":[{"Street":"jalan1"},{"Street":"jalan2"}]}`
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
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Address)

}