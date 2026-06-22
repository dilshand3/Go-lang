package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().Format("02/01/2006 15:04:05")
	fmt.Println(currentTime)

	formattedTime := "2026-06-22T06:35:50Z"
	parseTime, _ := time.Parse(time.UnixDate, formattedTime)
	fmt.Println(parseTime)
}
