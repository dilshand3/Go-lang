package main

import "fmt"

func updateByRefrence(num *int) {
	*num = *num + 10
}

func main() {
	num := 10

	fmt.Printf("This is value %d without modification", num)

	updateByRefrence(&num)

	fmt.Printf("This is value %d after modification", num)
}
