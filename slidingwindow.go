package main

import "fmt"

func subSums(nums []int, left int, right int) int {
	sum := 0
	sub := nums[left : right+1]
	for _, num := range sub {
		sum += num
	}
	return sum
}

func slidWin() {
	left := 0
	currSum := 0
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	c := 0
	for right, num := range nums {
		currSum += num
		for left <= right && currSum > goal {
			currSum -= nums[left]
			left++
		}
		if subSums(nums, left, right) == goal {
			c++
		}
	}
	fmt.Println(c)
}
