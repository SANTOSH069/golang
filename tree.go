package main

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func Searchtree(root *Tree, key int) *Tree {
	for root != nil {
		if root.val == key {
			return root
		}
		if root.val < key {
			root = root.right
		} else {
			root = root.left
		}
	}
	return nil
}
