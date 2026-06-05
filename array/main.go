package main

import "fmt"

var arr = [3]int{1, 2, 3}

func main() {
	fruitArr := [3]string{"apple", "banana", "orange"}

	fmt.Println("This is my fruit: ", fruitArr);
	fmt.Println(len(fruitArr))
}
