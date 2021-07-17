package main

import (
	"fmt"
)

func main() {
	// That's a implementation if dynamic arrays, but that's bad, because create a Amortized time problem  https://medium.com/vendasta/golang-the-time-complexity-of-append-2177dcfb6bad
	var slice []int
	value :=0

	for i := 0; i < 5; i++ {
		fmt.Println("Ingrese un valor")

		_, err := fmt.Scanf("%d\n", &value)
		if err != nil {
			fmt.Println("error")
			return
		}

		slice = append(slice, value)
	}

	// We see the capacity is more than begin declaration, that's the problem called Amortized time
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	// We can create arrays wit the make sentence, that create an array with len an capacity defined
	sliceMake := make([]int, 5)

	for i := 0; i < 4; i++ {
		fmt.Println("Ingrese un valor")

		_, err := fmt.Scanf("%d\n", &value)
		if err != nil {
			fmt.Println("error")
			return
		}

		sliceMake[i] = value
	}

	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake))
}