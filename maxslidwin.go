package main

import (
	"fmt"
	"math"
)

func maxEle(arr []int) int {
	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func maxslidWin() {
	nums := []int{}
	k := 3
	left := 0
	res := []int{}
	arr := []int{}
	for right := 0; right < len(nums); right++ {
		arr = append(arr, nums[right])
		for right-left+1 > k {
			arr = arr[left+1:]
			left++
		}
		if len(arr) == k {
			res = append(res, maxEle(arr))
		}
	}
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])
	}
}
