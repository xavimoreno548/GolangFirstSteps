package main

import "fmt"

func main() {

	// The Pointers are elements containing memory address
	a := 25
	fmt.Println("Value of a: ", a)
	fmt.Println("Memory address of a: ", &a)

	// We're defining a pointer named b, b contains the address of a
	b := &a

	fmt.Println("Value of b: ", b)

	// We are change de value of b
	*b = 50

	fmt.Println("Value of b: ", *b)
	fmt.Println("Value of a: ", a)

}
