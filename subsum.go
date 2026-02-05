package main

import "fmt"

func subSum(nums []int, left int, right int) int {
	sum := 0
	sub := nums[left : right+1]
	for _, num := range sub {
		sum += num
	}
	return sum
}

func subarr() {
	left := 0
	currSum := 0
	k := 6
	conSub := false
	nums := []int{23, 2, 4, 6, 7}
	for right := 0; right < len(nums); right++ {
		currSum += nums[right]
		for right-left+1 < 2 {
			currSum -= nums[left]
			left++
		}
		if (right-left+1) >= 2 && (subSum(nums, left, right)%k == 0) {
			conSub = true
		}
		fmt.Println(conSub)
	}
}
