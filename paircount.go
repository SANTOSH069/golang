package main

import (
	"fmt"
	"log"
)

func isValid(sub []int, k int) bool {
	return sub[0]+sub[1] > k
}

func PairCount([]int) int {
	nums := []int{1, 2, 3, 4, 5}
	count := 0
	var k int
	_, err := fmt.Scanln(&k)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sub := nums[i : j+1]
			if len(sub) == 2 && j > i {
				if isValid(sub, k) {
					count++
				}
			}
		}
	}
	return count
}
