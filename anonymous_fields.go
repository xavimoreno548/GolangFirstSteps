package main

import "fmt"

// Person struct with information of person
type Person struct {
	name string
}

// UserOne struct with all information of Person
type UserOne struct {
	Person
}

func main() {
	userOne := UserOne{Person{"Charles"}}
	fmt.Println(userOne.Person.name)
}
