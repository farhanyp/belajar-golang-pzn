package main

import "fmt"

func main() {

	// Membuat slice
	names := [...]string{"farhan", "yudha", "pratama", "yp1", "yp2", "yp3"}

	// Mengambil array dari index 4 - sebelum index 6
	slice1 := names[4:6]
	fmt.Println(slice1)

	// Mengambil array dari sebelum index 6
	slice2 := names[:6]
	fmt.Println(slice2)

	// Mengambil array dari index 4 sampai terakhir
	slice3 := names[4:]
	fmt.Println(slice3)

	// Mengambil seluruh array
	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println("////////////////////")
	// function slice
	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumamt", "sabtu", "minggu",}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(len(daySlice1))
	fmt.Println(cap(daySlice1))

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice2)

	fmt.Println("////////////////////")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "slice1"
	newSlice[1] = "slice2"

	// kalau mau nambah slice harus dengan append
	newSlice2 := append(newSlice, "slice3")
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// salin slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)



}