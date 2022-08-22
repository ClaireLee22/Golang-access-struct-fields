package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	user := User{name: "John", age: 30}

	// use dot with a struct
	fmt.Printf("name: %s; age: %d\n", user.name, user.age)
}

/* output
name: John; age: 30
*/
