package main

import (
	"fmt"
)

func Pointers() {
	num := 42
	val := *&num
	fmt.Println(&num)
	fmt.Println(val)

}
