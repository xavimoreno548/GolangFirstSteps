package main

import "fmt"

func main() {
	// 	Standard definition for array
	/* 	In Go the vars are initialized with de default value of each type of data
		int 				--> 0
		string 				--> ""
		bool				--> false
		float32 / float64	--> 0
	 */
	var arrayOfInt [5]int
	var arrayOfString [5]string
	var arrayOfBool [5]bool
	var arrayOfFloat [5]float32
	fmt.Println("Array of Int: ", arrayOfInt, "\n",
					"Array of String: ", arrayOfString, "\n",
					"Array of Bool: ", arrayOfBool, "\n",
					"Array of Float: ",arrayOfFloat)

	// Iterations of arrays, usign foreach form of for
	for _, v := range arrayOfInt {
		fmt.Println(v)
	}
}
