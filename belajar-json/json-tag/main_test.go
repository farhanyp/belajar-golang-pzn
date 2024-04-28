package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestEncodingJson(t *testing.T) {

	product := Product{
		Id: "1",
		Name: "yp",
		ImageUrl: "tes",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestDicodingJsonArray(t *testing.T) {

	jsonString := `{"id":"1","name":"yp","image_url":"tes"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)

}