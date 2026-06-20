package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}

	for index, values := range number {
		fmt.Printf("%d value index is %d\n",values,index)
	}
}
