package main

import "fmt"

// membuat function dengan parameter dan juga return value
func sayHello(name string) string {
	return "Hello " + name
}

// membuat multiple return value
func fullname() (string, string, string) {
	return "Farhan ", "Yudha", "Pratama"
}

// membuat named return value
func getFullname() (firstname, middlename, lastname string) {
	firstname = "Farhan"
	middlename = "Yudha"
	lastname = "Pratama"

	return firstname, middlename, lastname
}

// membuat varadic function
func getAllNames(names ...string) {

	for _,name := range names{
		fmt.Println(name)
	}
}	

// function as value
func functionAsValue(names string) string {
	return names
}

// function as parameter
func getNameByFilter1(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func FilterName1(name string) string {
	if name == "anjing"{
		return "..."
	}else{
		return name
	}
}

// function as parameter with Type
type Filter func(string) string

func getNameByFilter2(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func FilterName2(name string) string {
	if name == "anjing"{
		return "..."
	}else{
		return name
	}
}


// recursive function
func Factorial(number int) int {
	if number == 1{
		return 1
	}else{
		return number * Factorial(number - 1)
	}
}

func main() {

	// hello := sayHello("farhan")
	// firstname, middlename, lastname := getFullname()
	// fmt.Println(firstname, middlename, lastname)

	// getAllNames("farhan", "yudha", "pratama")
	// var names = []string{"yp1", "yp2", "yp3"}
	// getAllNames(names...)

	// var name = functionAsValue
	// println(name("farhan"))

	// getNameByFilter1("farhan", FilterName1)
	// getNameByFilter2("maman", FilterName2)


	// Anonymous function
	// getNameByFilter2("racing", func(name string) string {
	// 	if name == "anjing"{
	// 		return "..."
	// 	}else{
	// 		return name
	// 	}
	// })


	// Closure
	// Closure itu berarti mengakses data yang berada di sekitar scopenya
	// var counter int = 0
	// increment := func()  {
	// 	fmt.Println("Increment")
	// 	counter++
	// }

	// increment()
	// increment()
	// increment()
	// fmt.Println("Counter", counter)
}