package main

import (
	"fmt"
)

func primeNos(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primerange() {
	res := []int{}
	num := 10
	for i := 2; i < num; i++ {
		if primeNos(i) {
			res = append(res, i)
		}
	}
	fmt.Println(res)
}
