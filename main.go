package main

import (
	"fmt"
	"strings"
)

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	res := strings.Fields(paragraph)
	words := []string{}
	for _, word := range res {
		words = append(words, word)
	}
	for _, str := range words {
		fmt.Print(str + " ")
	}
}
