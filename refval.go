package main

import "fmt"

func AddVal(a *int, b *int) {
	*a *= 100
	*b *= 200
}

func refval() {
	var a int = 10
	var b int = 20
	add := &a
	fmt.Println("The value of a before changing: %d\n", a)
	fmt.Println("The value of a before changing: %d\n", b)

	AddVal(&a, &b)
	fmt.Println("The address of the varibale a is: %d", add)
	fmt.Println("The value of a after change is: %d", a)
	fmt.Println("The value of b after change is: %d", b)

}
