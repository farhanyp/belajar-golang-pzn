package main

import "fmt"

// Cara membuat struct
// cara 1
type User struct {
	Name, Addres string
	age          int
}

func (user User) getUser(){
	fmt.Println("ini method getuser adalah:", user.Name)
}


func main() {


	var yp User
	yp.Name = "Farhan"
	yp.Addres = "Jauh"
	yp.age = 20

	// Cara 2
	farhan := User{
		Name:   "Farhan",
		Addres: "Jauh",
		age:    20,
	}

	// Cara 3
	farhanyp := User{"farhan", "user", 20}

	fmt.Println(yp)
	fmt.Println(farhan)
	fmt.Println(farhanyp)

	farhanyp.getUser()

}