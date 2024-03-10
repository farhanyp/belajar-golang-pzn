package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile(`y([a-z])p`)

	fmt.Println(regex.MatchString("yap"))
	fmt.Println(regex.MatchString("yep"))
	fmt.Println(regex.MatchString("yop"))
	fmt.Println(regex.MatchString("ybp"))
	fmt.Println(regex.MatchString("ycp"))

	fmt.Println(regex.FindAllString("ycp yap yah yux",5))

}