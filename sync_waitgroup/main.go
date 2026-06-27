package main

import (
	"fmt"
	"sync"
)

func worker(i int, wq *sync.WaitGroup) {
	fmt.Printf("This is %d worker start\n", i)
	fmt.Printf("This is %d worker end\n", i)

	defer wq.Done()
}

func main() {

	var wq sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wq.Add(1)
		go worker(i, &wq)
	}
	wq.Wait()

	fmt.Println("This is full program exit")
}
