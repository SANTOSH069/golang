package main

import "fmt"

func printNum(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	ch <- 1
}

func printAlpha(ch chan int) {
	for i := 'A'; i < 'F'; i++ {
		fmt.Println(string(i))
	}
	ch <- 2
}

func Chann() {

	ch := make(chan int)
	go printNum(ch)
	go printAlpha(ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y)
}
