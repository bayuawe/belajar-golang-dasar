package main

import "fmt"

func main() {

	// var name string

	// name = "Bayu Aryandi"
	// fmt.Println(name)

	// name = "Bayu Wijaya"
	// fmt.Println(name)

	//bisa juga langsung seperti ini

	// var name = "Bayu Aryandi" // golang bisa tau tipe data yang di gunakan adalah string
	// fmt.Println(name)

	// name = "Bayu Wijaya"
	// fmt.Println(name)

	name := "Bayu Aryandi" //bisa juga seperti ini tapi hanya saat pertama kali deklarasi variable
	fmt.Println(name)

	name = "Bayu Wijaya"
	fmt.Println(name)

	var (
		firstName = "Bayu Aryandi"
		lastName  = "Wijaya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
