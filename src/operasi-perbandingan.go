package main

import "fmt"

func main() {
	var nama1 = "rendi"
	var nama2 = "putra"

	var result bool = nama1 == nama2
	fmt.Println(result)

	var angka1 = 10
	var angka2 = 9

	result = angka1 > angka2 // angka 1 lebih besar dari pada angka 2
	fmt.Println(result)
	fmt.Println(angka1 != angka2) // true
}
