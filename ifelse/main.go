package main

import (
	"fmt"
)

func main(){
	a := 11;
	b := 11;

	if a > b {
		fmt.Println("a is greater then b");
	}else if b > a {
		fmt.Println("b is greater then a")
	}else {
		fmt.Println("a is equal to b")
	}
}