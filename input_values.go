package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Obtención de valores con entrada estandar
	var number int

	fmt.Println("Ingrese un valor entero")
	fmt.Scanf("%d\n", number)

	fmt.Printf("El número es %d\n", number)

	// Obtencion de valores con bufio
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese su nombre")
	userName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(" El nombre es " + userName)
	}

}