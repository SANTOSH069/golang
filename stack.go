package main

import (
	"fmt"
)

type Stack struct {
	st []int
}

func (s *Stack) Push(item int) {
	s.st = append(s.st, item)
	fmt.Println("Item Pushed")
}

func (s *Stack) isEmpty() bool {
	if len(s.st) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() {
	if s.isEmpty() {
		fmt.Println("Nothing to pop!")
	}
	last := len(s.st) - 1
	s.st = s.st[:last]
	fmt.Println("Element Deleted")
}

func (s *Stack) Peek() int {
	res := s.st[0]
	return res
}
