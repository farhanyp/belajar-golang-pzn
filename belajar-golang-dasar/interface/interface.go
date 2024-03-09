package main

import "fmt"

type Habits interface {
	GetName() string
}

type User struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHello(habit Habits) {
	fmt.Println(habit.GetName())
}

func (user User) GetName() string {
	return user.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	yp := User{"farhanyp"}
	sayHello(yp)

	animal := Animal{"anjing"}
	sayHello(animal)

}