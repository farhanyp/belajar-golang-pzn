package main

import "fmt"

// nil merupakan pengertian dari object yang belum terinisialisai
// nil hanya bisa digunakan pada interface, function, map, slice, pointer dan channel
func getDataMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	data := getDataMap("tes")

	if data == nil {
		fmt.Println("data kosong")
	}else{
		fmt.Println(data["name"])
	}

}