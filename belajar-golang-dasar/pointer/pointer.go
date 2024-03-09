package main

import "fmt"

type Bio struct {
	name, Address, Hobby string
}

// Contoh function dengan parameter pointer
func ChangeAddressToIndonesia(address *Bio) {
	address.Address = "Indonesia"
}

// Contoh pointer pada method struct
func (bio *Bio) changeNameToYp() {
	bio.name = "Yp"
}

func main() {

	// Hanya mengcopy value dari name 1 ke name 2
	// var name1 Bio = Bio{"farhan", "jauh1", "belajar"}
	// var name2 Bio = name1
	// name2.name = "Yp"

	// Menggunakan pointer
	var name3 Bio = Bio{"farhan", "jauh1", "belajar"}
	// name4 mereferensi ke name3
	var name4 *Bio = &name3
	// Disini name4 merubah isi dari reference name 3, jadinya tipe data dari name4 yaitu pointer
	name4.name = "Yp Besar"

	// Disini name 4 membuat reference baru, karena name 4 bertipe data pointer jadi harus di tambah & di depannya
	name4 = &Bio{"Yp Kecil", "jauh1", "belajar"}

	// Menggunakan Arterisk Operator
	var name5 Bio = Bio{"farhan", "jauh1", "belajar"}
	var name6 *Bio = &name5

	// Disini name5 sudah mereference ke name6
	*name6 = Bio{"Yp Kecil", "jauh1", "belajar"}
	name6.name = "Yp berubah"

	// Function name()
	// untuk membuat pointer baru tapi kosong
	// user1 := new(Bio)

	// Mencoba Function dengan parameter pointer
	// cara 1
	// user1 := &Bio{Address: "Amerika"}
	// ChangeAddressToIndonesia(user1)
	// fmt.Println(user1)

	// cara 2
	// user1 := Bio{Address: "Amerika"}
	// ChangeAddressToIndonesia(&user1)
	// fmt.Println(user1)


	// Mencoba method struct dengan parameter pointer
	user1 := Bio{name: "Lontong"}
	user1.changeNameToYp()

	fmt.Println(user1)
}