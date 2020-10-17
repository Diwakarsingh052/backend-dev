package main

import "fmt"

// Q. Pointer on pointer?

func pointers() int {
	a := 10
	var p1 *int
	p1 = &a

	// pointer to pointer
	var p2 **int
	p2 = &p1

	return **p2
}

func main() {
	fmt.Println(pointers())
}
