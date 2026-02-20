package main

import "fmt"

func unqFreq() {
	freq := [26]int{}

	str := "abcde"
	str2 := "abcd"
	res := str + str2

	for _, ch := range res {
		freq[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq[i] == 1 {
			fmt.Printf("Unique character: %c\n", rune(i+'a'))
		}
	}
}
