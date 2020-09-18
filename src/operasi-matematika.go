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

	// Augmented Assignments
	// a += 10  ==  a = a + 10
	fmt.Println("---------")
	a += 10        //a = 10 + 10
	fmt.Println(a) //20

	a = 10
	a -= 10        //a = 10 - 10
	fmt.Println(a) //0

	a = 10
	a *= 10        //a = 10 x 10
	fmt.Println(a) //100

	a = 10
	a /= 10        //a = 10 / 10
	fmt.Println(a) //1

	// Unary Operator
	// a++  ==  a = a + 1
	fmt.Println("---------")
	a = 10
	a++            //a = 10 + 1
	fmt.Println(a) //11

	a = 10
	a--            //a = 10 - 1
	fmt.Println(a) //9

}
