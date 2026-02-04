package main

import ()

func binSub() {
	nums := []int{1, 0, 1, 0, 1}
	c := 0
	hash := make(map[int]int)
	k := 2
	hash[0] = 1
	prefSum := 0
	for i := 0; i < len(nums); i++ {
		prefSum += nums[i]
		goal := prefSum - k
		if freq, exists := hash[goal]; exists {
			c += freq
		}
		hash[prefSum]++
	}

}
