package main

import (
	"fmt"
	"strings"
)

func freqWord() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."

	words := strings.Fields(paragraph)
	fmt.Print(words)
}
