package main

import "fmt"

func main() {

	// Membuat map
	// var name map[string]string = map[string]string{}
	// name := map[string]string{
	// 	"first": "farhan",
	// }
	name := make(map[string]string)

	name["first"] = "farhan"
	name["midde"] = "yudha"
	name["last"] = "pratama"
	name["wrong"] = "wrong"

	delete(name, "wrong")
	fmt.Println(name)
	fmt.Println(name["first"])
	fmt.Println(len(name))
	fmt.Println(len(name["middle"]))

}