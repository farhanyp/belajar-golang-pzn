package main

import (
	"container/list"
	"fmt"
)

func main() {

	var data *list.List = list.New()

	data.PushBack("Farhan")
	data.PushBack("Yudha")
	data.PushBack("Pratama")

	// cara melihat value secara manual
	// var head *list.Element =  data.Front()
	// fmt.Println(head.Value)

	// var next *list.Element =  head.Next()
	// fmt.Println(next.Value)

	// next =  next.Next()
	// fmt.Println(next.Value)
	
	// cara melihat value secara perulangan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}