package main

import "fmt"

func main(){
	num := 3

	a := &num

	fmt.Println(*a)
	fmt.Println(a)
}