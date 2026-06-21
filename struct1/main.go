package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	user1 := Person{
		firstName: "dilshan",
		lastName:  "hansal",
		age: 12,
	}

	fmt.Println(user1)
}
