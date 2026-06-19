package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter user name")

	// var name string;

	// fmt.Scan(&name);

	// fmt.Println("This is username:",name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println(name)
}
