package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	student1 := Person{
		Name:    "dilshan",
		Age:     18,
		IsAdult: true,
	}

	student1JsonData, student1Error := json.Marshal(student1)

	if student1Error != nil {
		fmt.Println(student1Error)
	}

	fmt.Println(string(student1JsonData))

	var parseData Person

	err2 := json.Unmarshal(student1JsonData, &parseData)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(parseData)
}
