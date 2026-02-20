package main

import "fmt"

func revStr(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func unqChar() {
	str := "abcde"
	str2 := "abcd"
	s := str + revStr(str2)
	arr := []rune(s)
	var xor rune = 0
	for _, ch := range arr {
		xor ^= ch
	}
	fmt.Println(xor)
}
