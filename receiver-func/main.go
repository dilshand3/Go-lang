package main

import "fmt"

type Car struct {
	Color string
}

func (p Car) sayHello() string {
	return p.Color + "car"
}

func main() {
	c := Car{Color: "red"}

	data := c.sayHello()

	fmt.Println(data)
}
