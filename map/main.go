package main

import "fmt"


func main(){
	studentMap := make(map[string]int);

	studentMap["dilshan"] = 18;
	studentMap["nitika"] = 18;
	studentMap["nikhil"] = 18;
	studentMap["vikas"] = 19;

	fmt.Println(studentMap["vikas"])
}