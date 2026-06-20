package main

import (
	"fmt"
)

func main() {
	playersMap := make(map[string]int)

	playersMap["dilshan"] = 12
	playersMap["nikhil"] = 11
	playersMap["vikas"] = 10

	zaheer, exists := playersMap["zaheer"]

	fmt.Println(zaheer,exists)
}
