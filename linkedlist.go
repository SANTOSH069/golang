package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertNode(data int) {
	newNode := &Node{val: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *LinkedList) ContainsNode(k int) bool {
	curr := l.head
	for curr != nil {
		if curr.val == k {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *LinkedList) DeleteAtSpecific(pos int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if pos == 1 {
		l.head = l.head.next
		return
	}
	curr := l.head
	for i := 1; i < pos-1 && curr.next != nil; i++ {
		curr = curr.next
	}
	if curr.next == nil {
		fmt.Println("Invalid position")
		return
	}
	curr.next = curr.next.next
}

func (l *LinkedList) DisplayNode() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) DeleteFirst() {
	if l.head == nil {
		fmt.Println("Nothing to delete!")
	}
	curr := l.head
	l.head = curr.next
	fmt.Println("Node Deleted !")
}
