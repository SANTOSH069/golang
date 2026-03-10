package main

// "fmt"

func minEle(sub []int) int {
	min := 12345678
	for _, num := range sub {
		if min > num {
			min = num
		}
	}
	return min
}

func isEle(sub []int) bool {
	Min := minEle(sub)
	return Min == sub[0]
}

func minSub(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sub := nums[i : j+1]
			if isEle(sub) {
				count++
			}
		}
	}
	return count
}
