package main

import (
	"fmt"
)

func main(){
	fmt.Println("Enter user name");

	var name string;

	fmt.Scan(&name);

	fmt.Println("This is username:",name)
}