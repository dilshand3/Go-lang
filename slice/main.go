package main

import "fmt"

var slice = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(slice)

	makeSlice := make([]int, 0, 10)
	fmt.Println(makeSlice)
}
