package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	data, err2 := ioutil.ReadAll(res.Body)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(string(data))
}
