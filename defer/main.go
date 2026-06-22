package main

import (
	"fmt"
)

func main(){
	fmt.Println("This is line no.1")
	defer fmt.Println("This is line no.2")
	fmt.Println("This is line no.3")
}