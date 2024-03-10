package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Date())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Year())


	// Duration
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration : %d\n", duration3)
}