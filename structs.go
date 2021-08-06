package main

import "fmt"

// User struct to represent users, contains a name and age
type User struct {
	name string
	age int
}

// Function to set the name of user, that function belongs to the User struct
func (user *User) setName(name string)  {
	user.name = name
}

// Function to get the name of user, that function belongs to the User struct
func (user User) getName() string {
	return user.name
}

func main() {
	// We created a variable of type User, and fill the name and age
	userOne := User{name: "Xavi", age: 28}
	fmt.Println(userOne)
}
