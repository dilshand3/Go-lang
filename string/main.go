package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hello,world"
	data := strings.Split(str1, ",")
	fmt.Println(data)

	str2 := "   hello dilshan    "
	fmt.Println(strings.TrimSpace(str2))

	str3 := "dilshan"
	str4 := "hansal"
	fmt.Println(strings.Join([]string{str3, str4}, " "))
}
