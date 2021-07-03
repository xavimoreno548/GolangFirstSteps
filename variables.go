package main

import (
"fmt"
"strconv"
)

func main() {
	// Decalracion de variables
	var number int
	var name string
	var weight float32
	var flag bool

	number = 5
	name = "Xavi"
	weight = 82.5
	flag = true

	// Impresión de valores
	fmt.Println("Este es el nombre: " + name)

	// Impresión de valores con verbos
	fmt.Printf("Este es el número: %d\n", number)
	fmt.Printf("Este es el peso %f\n", weight)
	fmt.Printf("Esta es la bandera %v\n", flag)

	// Conversión de entero a string
	fmt.Println("Este es el número: " + strconv.Itoa(number))
}