package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()
}
