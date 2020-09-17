package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRendi NoKTP = "21421120001"
	fmt.Println(ktpRendi)
	fmt.Println(NoKTP("5234134002350001"))

}
