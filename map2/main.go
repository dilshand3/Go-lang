package main

import (
	"fmt"
)

func main() {
	studentMap := make(map[string]int)

	studentMap["dilshan"] = 98

	fmt.Println(studentMap["dilshan"])
}
