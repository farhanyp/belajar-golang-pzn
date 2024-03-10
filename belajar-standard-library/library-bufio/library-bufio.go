package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Membaca inputan
	input := strings.NewReader("This is long string")

	reader := bufio.NewReader(input)

	for{
		line, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		fmt.Println(string(line))
	}

	// Menulis Inputan
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello world\n")

	writer.Flush()

	// Menulis dari terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write something: ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("You write:", text)

}