package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodingJson(t *testing.T) {

	product := map[string]interface{}{
		"id": "1",
		"name": "yp",
		"image_url": "tes",
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

	var product map[string]interface{}

	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)

}