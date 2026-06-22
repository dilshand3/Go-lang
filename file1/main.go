package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content := "this is my file"
	_, err2 := io.WriteString(file, content)

	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
