package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	var c = a + b
	fmt.Println(a, " + ", b, " = ", c)

	c = a - b
	fmt.Println(a, " - ", b, " = ", c)

	c = a / b
	fmt.Println(a, " / ", b, " = ", c)

	c = a * b
	fmt.Println(a, " X ", b, " = ", c)

	// Augmented Assignments Command

	a = a + 10     //a = 10 + 10
	fmt.Println(a) //20

}
