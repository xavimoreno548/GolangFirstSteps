package main

import "fmt"

func main() {
	// Classic C/C++ for
	var i int
	for i = 0; i<10; i++ {
		fmt.Printf("Classic for - i: %d\n", i)
	}

	// While form
	i = 0
	for i<10 {
		fmt.Printf("While form for - i: %d\n", i)
		i++
	}


	x := []int{0, 0, 0}

	for index, y := range x {
		fmt.Printf("indice %d\n", index)
		fmt.Printf("Valor %d\n", y)

	}
}