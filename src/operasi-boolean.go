package main

import "fmt"

func main() {
	var nilaiIPA = 90
	var nilaiIPS = 60

	var kkmIPA bool = nilaiIPA > 73
	var kkmIPS bool = nilaiIPS > 78

	var lulus bool = kkmIPA && kkmIPS // and

	fmt.Println(lulus)

	/*
		## 	Operator Logika

			&&	:	And
			||	:	Or
			!	:	Not
	*/

	// cara kedua
	fmt.Println(nilaiIPA >= 74 && nilaiIPS >= 78)
}
