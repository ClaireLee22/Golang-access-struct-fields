package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	userPointer := &User{name: "John", age: 30}

	// use dot with a struct pointer
	fmt.Printf("name: %s; age: %d\n", userPointer.name, userPointer.age)

	// parenthesize the struct pointer and then access field with a dot operator
	fmt.Printf("name: %s; age: %d\n", (*userPointer).name, (*userPointer).age)

	// // error: invalid operation: cannot indirect userPointer.name (variable of type string)
	// fmt.Printf("name: %s; age: %d\n", *userPointer.name, *userPointer.age)
}

/* output
name: John; age: 30
name: John; age: 30
*/
