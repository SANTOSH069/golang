package main

import (
	"fmt"
)

func issubstr() bool {
	s := "hello"
	t := "ell"
	for i := 0; i <= len(s)-len(t); i++ {
		if t == s[i:i+len(t)] {
			return true
		}
	}
	return false
}
