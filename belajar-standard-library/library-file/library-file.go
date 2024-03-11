package main

import (
	"bufio"
	"io"
	"os"
)

// func createNewFile(name string, message string) error {
// 	file, err := os.OpenFile(name, os.O_CREATE | os.O_WRONLY, 0)
// 	if err != nil{
// 		return err
// 	}
// 	defer file.Close()

// 	file.WriteString(message + "\n")
// 	return nil
// }


// func readFile(name string) (string, error){
// 	file, err := os.OpenFile(name, os.O_RDONLY, 0)
// 	if err != nil{
// 		return "", err
// 	}
// 	defer file.Close()

// 	reader := bufio.NewReader(file)
// 	var message string

// 	for{
// 		line, _, err := reader.ReadLine()
// 		if err == io.EOF{
// 			break
// 		}

// 		message += string(line)
// 	}

// 	return message, nil
// }


func addToFile(name string, message string) error{
	file, err := os.OpenFile(name, os.O_RDWR | os.O_APPEND, 0)
	if err != nil{
		return err
	}
	defer file.Close()

	file.WriteString(message)

	return nil
}

func main() {

	// _ = createNewFile("library-file/tes.log", "ini adalah log")
	// message, _ := readFile("library-file/tes.log")
	
	_ = addToFile("library-file/tes.log", "Menambah Log")


}