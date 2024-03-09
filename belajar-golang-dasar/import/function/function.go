package function

import "fmt"

var version string = "1.1.1" // tidak bisa diakses diluar package
var Application string = "1.1.1" // bisa diakses diluar package

func Hello() {
    // Access Modifier bergantung besar dan kecilnya huruf
    fmt.Println("Hello from import.go")
}