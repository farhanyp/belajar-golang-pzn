package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}){
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJson(t *testing.T) {

	LogJson("data")
	LogJson(1)
	LogJson([]string{"farhan", "yudha", "pratama"})

}