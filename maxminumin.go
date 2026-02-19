package main

import (
	"fmt"
	"slices"
)

func minVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func finMinMax(nums []int, idx int) int {
	arr := nums[:]
	arr[idx] *= 2
	min := slices.Min(arr)
	max := slices.Max(arr)
	return max - min
}

func minMax() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	nums := []int{}
	i := 0
	for i <= n {
		var x int
		fmt.Scanf("%d", &x)
		nums = append(nums, x)
		i++
	}
	Min := 12345678
	for i := 0; i < len(nums); i++ {
		Min = minVal(Min, finMinMax(nums, i))
	}
	fmt.Println(Min)

}
