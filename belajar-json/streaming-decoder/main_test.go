package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct{
	Firstname, Middlename, Lastname string
}

func TestDecoderJson(t *testing.T) {

	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestEncoderJson(t *testing.T) {

	writer, _ := os.Create("sample-out.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		Firstname: "farhan",
		Middlename: "yudha",
		Lastname: "pratama",
	}

	encoder.Encode(customer)

}

// func TestDicodingJsonArray(t *testing.T) {

// 	jsonString := `{"id":"1","name":"yp","image_url":"tes"}`
// 	jsonBytes := []byte(jsonString)

// 	var product map[string]interface{}

// 	err := json.Unmarshal(jsonBytes, &product)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(product)

// }