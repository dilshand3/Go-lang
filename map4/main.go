package main

import (
	"fmt"
)

func main() {
	studentsMap := map[string]int{
		"dilshan": 1,
		"zaheer":  2,
		"nikhil":  3,
	}

	fmt.Println(studentsMap)

	dilshanKey, dilshanExists := studentsMap["dilshan"]

	fmt.Println(dilshanKey, dilshanExists)
}
