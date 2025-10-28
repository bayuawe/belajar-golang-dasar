package main

import "fmt"

func main() {

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//variable tidak sanggup menampung data
	fmt.Println(nilai16) // nilai mengalami overflow sehingga reset kembali ke angka paling kecil

	//conversion data (2)

	var name = "Bayu Aryandi Wijaya"
	var e = name[0]         // mengambil index ke 0 pada variable name
	var eString = string(e) //mengubah tipe data uint8 ke string

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
