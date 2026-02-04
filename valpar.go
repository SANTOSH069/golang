package main

import "fmt"

func valPar() {
	s := "()"
	open := 0
	close := 0
	arr := []rune(s)
	for _, ch := range arr {
		if ch == '(' || ch == '{' || ch == '[' {
			open++
		}
		if open > 0 {
			open--
		}
		close++
	}
	if close > 0 {
		fmt.Println(false)
	}
	fmt.Println(true)

}
