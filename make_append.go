package main

import "fmt"

func main() {

	// We use that notation, for create a slice with the same capacity and length
	slice := make([]int, 3)
	fmt.Println("Slice with same capacity and length")
	fmt.Println(slice)

	// We use that notation for create a slice with different capacity and length
	slice = make([]int, 3, 4)
	fmt.Println("Slice with different capacity and length")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// We use append to add values to end of slice
	// That is cool with we don't exceed the capacity
	slice = append(slice, 5)
	fmt.Println(slice)

	// When we exceed the capacity, we have a problem, because the slice overflow and we create a Amortized time problem
	slice = append(slice, 5)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}
