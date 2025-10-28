package main

import "fmt"

func main() {
	// const firstName string = "Bayu Aryandi"
	// const lastName = "Wijaya"

	// //Tidak bisa di assign ulang jika variable menggunakan constant
	// firstName = "Ica" //cannot assign to firstName (neither addressable nor a map index expression)
	// lastName = "Jaya" //cannot assign to lastName (neither addressable nor a map index expression)

	const (
		firstName string = "Bayu Aryandi"
		lastName         = "Wijaya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
