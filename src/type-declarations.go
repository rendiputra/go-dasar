package main

import "fmt"

func main() {

	// Membuat alias variable:

	// variabel 'string' diubah menjadi 'NoKTP'
	type NoKTP string

	// ketika menambah variable yang seharusnya 'var ktpRendi string = "21421120001"' diubah menjadi:
	var ktpRendi NoKTP = "21421120001"
	fmt.Println(ktpRendi)
	fmt.Println(NoKTP("5234134002350001"))

}
