package main

import (
	"fmt"
	"sync"
)

func printNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	ch <- 1
}

func printAlpha(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i < 'F'; i++ {
		fmt.Println(string(i))
	}
	ch <- 2
}

func Chann() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go printNum(ch, &wg)
	wg.Add(2)
	go printAlpha(ch, &wg)

	x, y := <-ch, <-ch
	fmt.Println(x, y)
	wg.Wait()
}
