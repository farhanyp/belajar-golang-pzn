package main

import (
	"fmt"
	"sort"
)

type Users struct {
	Name string
	Age  int
}

type userSlice []Users

// Membuat contract untuk sort
func (userSlice userSlice) Len() int {
	return len(userSlice)
}

func (userSlice userSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice userSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

	users := []Users{
		{"yp1", 15},
		{"yp2", 14},
		{"yp3", 12},
		{"yp4", 18},
		{"yp5", 9},
	}

	sort.Sort(userSlice(users))

	fmt.Println(users)

}