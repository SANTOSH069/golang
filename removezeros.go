package main

import "fmt"

func removezeroes() {
	nums := []int{0, 1, 0, 3, 12}
	res := []int{}
	zeroes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums = append(nums, nums[i])
		}
		if nums[i] == 0 {
			zeroes++
		}
	}
	for zeroes > 0 {
		res = append(res, 0)
		zeroes--
	}
	for _, i := range res {
		fmt.Print(i)
	}
}
