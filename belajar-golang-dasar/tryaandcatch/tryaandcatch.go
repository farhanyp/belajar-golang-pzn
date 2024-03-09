package main

import "fmt";

func endApp(){
	message := recover()
	fmt.Println("Terjadi error:",message)
	fmt.Println("End  App")
}

func startApp(error bool){
	defer endApp()

	fmt.Println("Start App")
	if error{
		panic("Error Cuy")
	}

}

func main(){
	// defer dibuat di awal function tetapi bakal keluar di akhir function
	startApp(true)

}