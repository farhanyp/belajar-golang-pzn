package main

func main() {

	// alias
	// noKTP bisa juga dibilang string
	type noKTP string

	var ktp1 noKTP = "111111"

	var nomor string = "222222"

	var ktp2 noKTP = noKTP(nomor)

}