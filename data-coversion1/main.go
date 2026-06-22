package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 12

	aString := strconv.Itoa(a)

	fmt.Printf("This is values %d and its type is %T\n", a, aString)

	b := "12"

	bInt,_ := strconv.Atoi(b)

	fmt.Printf("this is value %s and its type is %T\n", b, bInt)
}
