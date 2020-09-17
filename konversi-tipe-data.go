package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	// Apabila daya tampung yang tidak muat maka 'Value' akan berubah
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// Apabila daya tampung yang tidak muat maka 'Value' akan berubah
	fmt.Println(nilai8)

	fmt.Println("\n------------------------\n")

	var value32 int32 = 10
	var value64 int64 = int64(value32)

	// Apabila daya tampung yang muat maka 'Value' akan tetap
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)

	// Apabila daya tampung yang muat maka 'Value' akan tetap
	fmt.Println(value8)

	fmt.Println("\n------------------------\n")

	var nama = "Rendi Putra Pradana"
	var e byte = nama[0]
	var eString string = string(e)

	fmt.Println(nama)
	fmt.Println(eString)
}
