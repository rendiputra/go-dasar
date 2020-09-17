package main

import "fmt"

func main() {

	// Constant Wajib diberi Value diawal
	const nama = "Rendi Putra"
	const phi = 3.14
	const fibonnaci = "0 1 1 2 3 5 8 13.."

	// Multiple Deklarasi
	const (
		namaPanjang = "Rendi Putra Pradana"
		namaPendek  = "Rendi"
	)
	/*
		---------------
		-	note      -
		---------------

		Error : Deklarasi Ulang
		phi = 10

		tidak wajib digunakan
		fmt.Println(nama)

	*/

	fmt.Println(phi)
	fmt.Println(fibonnaci)
}
