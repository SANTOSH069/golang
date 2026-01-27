package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minDiff(arr []int) int {
	minDiff := int(^uint(0) >> 1)

	for i := 0; i < len(arr)-1; i++ {
		currDiff := abs(arr[i+1] - arr[i])
		if currDiff < minDiff {
			minDiff = currDiff
		}
	}
	return minDiff
}

func subArr(min int, arr []int) [][]int {
	var res [][]int

	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}

func minabs() {
	arr := []int{4, 2, 1, 3}

	sort.Ints(arr)

	min := minDiff(arr)
	res := subArr(min, arr)

	for _, pair := range res {
		fmt.Println(pair)
	}
}
