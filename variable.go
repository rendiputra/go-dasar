package main

import "fmt"

func main() {
	/*
		Variable harus digunakan, Jika tidak maka akan Error
	*/

	var name string

	name = "Rendi Putra Pradana"
	fmt.Println("Nama : ", name)

	name = "Rendi"
	fmt.Println("Nama Pendek : ", name, "\n\n----------------------\n")

	// Multiple Deklarasi
	var (
		namaLengkap = "Rendi Putra Pradana"
		namaPendek  = "Rendi"
	)

	fmt.Println("Nama Lengkap : ", namaLengkap)
	fmt.Println("Nama Pendek : ", namaPendek)

	var angka = 12
	fmt.Println("angka : ", angka)

	// Menggunakan int8
	var umur int8 = 18
	fmt.Println("umur : ", umur)

	// Tanpa var
	jurusan := "RPL"
	fmt.Println("Jurusan : ", jurusan)
}
