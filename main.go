package main

import (
	"fmt"
	"sync"
)

func display(str string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go display("This is Go-Routines", &wg)
	go display("Hello There", &wg)

	wg.Wait()
}
