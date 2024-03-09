package main

import "fmt"

type Bio struct {
	name, Address, Hobby string
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

	fmt.Println(name5)
	fmt.Println(name6)
}