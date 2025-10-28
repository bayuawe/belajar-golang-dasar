package main

import "fmt"

func main() {
	type NoKTP string

	var ktpBayu NoKTP = "111111111"

	var contoh string = "22222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpBayu)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}
