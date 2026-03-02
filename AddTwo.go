package main

import (
	"fmt"
	"testing"
)

func addTwo(a int, b int) int {
	return a + b
}

func testAdd(t *testing.T) {
	a := 2
	b := 3
	res := addTwo(a, b)
	want := 5
	if res != want {
		t.Errorf("Add(2, 3) = %d; want %d", res, want)
	}
	fmt.Println(res)
}
