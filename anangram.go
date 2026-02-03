package main

import (
	"fmt"
	"sort"
)

func isAnang() {
	s := "dog"
	t := "god"
	IsAnang := true
	if len(s) != len(t) {
		IsAnang = false
	}
	arr1 := []rune(s)
	arr2 := []rune(t)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	hash := make(map[rune]int)

	for _, ch := range arr1 {
		hash[ch]++
	}

	for _, ch := range arr2 {
		hash[ch]--
	}

	for _, val := range hash {
		if val != 0 {
			IsAnang = false
			break
		}
	}
	fmt.Println(IsAnang)

}
