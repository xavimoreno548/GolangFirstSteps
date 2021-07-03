package main

import "fmt"

func main() {
	// Condicionales
	x := 5
	y := 10

	if y > x {
		fmt.Printf("%d es mayor que %d", y,x)
	} else if x > y {
		fmt.Printf("%d es mayor que %d", x, y)
	} else {
		fmt.Printf("%d y %d son iguales", y,x)
	}

}