package main

import (
	"fmt"
	"time"
)

func sayHi() {
	fmt.Println("This is second line")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("This is third line")
}

func main() {
	fmt.Println("This is first line")
	go sayHi()

	fmt.Println("this is fouth line")
}
