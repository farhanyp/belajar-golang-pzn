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

func LogJson(){

	customer := Customer{
		Firstname: "Farhan",
		Middlename: "Yudha",
		Lastname: "Pratama",
		Age: 23,
		Ismarried: true,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJson(t *testing.T) {

	LogJson()

}