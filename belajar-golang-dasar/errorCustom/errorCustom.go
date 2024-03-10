package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id is empty"}
	}

	if id != "eko" {
		return &notFoundError{Message: "Data is Not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	// if err != nil {
	// 	if validationError, ok := err.(*validationError); ok {
	// 		fmt.Println("Validation error:", validationError.Message)
	// 	}else if notFoundError, ok := err.(*notFoundError); ok{
	// 		fmt.Println("Not Found error:", notFoundError.Message)
	// 	}else{
	// 		fmt.Println("unkown error:", err.Error())
	// 	}
	// }

	switch finalError := err.(type) {
	case *validationError:
		fmt.Println("Validation error:", finalError.Error())
	case *notFoundError:
		fmt.Println("Validation error:", finalError.Error())
	default:
		fmt.Println("unkown error:", finalError.Error())
	}
}