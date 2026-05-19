package main

import (
	"fmt"
	"runtime"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

}

func main() {
	fmt.Print("Memory Stats before: \n")

	s := make([]int, 10_00_000)

	for i := range s {
		s[i] = i
	}

}
