package main

import "fmt"

func main() {

	var nama [3]string // panjang(daya tampung) arraynya adalah 3, jika lebih maka akan error

	// index array dimulai dari 0
	nama[0] = "Rendi"
	nama[1] = "Putra"
	nama[2] = "Pradana"

	// Panggil index array yang dibutukan dengan memasukan indexnya
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])
	fmt.Println(nama[0], nama[1], nama[2])
	fmt.Println(nama, "\n---------------------------")

	// mengisi value array langsung saat deklarasi variable
	var buah = [4]string{
		"Apel",
		"Mangga",
		"Jambu",
	}

	fmt.Println(buah[0])
	fmt.Println(buah[1])
	fmt.Println(buah[2])
	fmt.Println(buah[0], buah[1], buah[2])
	fmt.Println(buah)

	// mengecek panjang(daya tampung) array
	fmt.Println(len(nama)) // 3
	fmt.Println(len(buah)) // 4

	/*
		Output:
		Rendi
		Putra
		Pradana
		Rendi Putra Pradana
		[Rendi Putra Pradana]
		---------------------------
		Apel
		Mangga
		Jambu
		Apel Mangga Jambu
		[Apel Mangga Jambu ]
		3
		4
	*/
}
