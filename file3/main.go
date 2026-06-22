package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Error while creating the file: ", err)
	}

	defer file.Close()

	content := "this is my data"
	_, err3 := file.WriteString(content)

	if err3 != nil {
		fmt.Println("this is error while writing file", err3)
	}

	data, err2 := os.ReadFile("example.txt")

	if err2 != nil {
		fmt.Println("can't read the file")
	}

	fmt.Println(string(data))
}
