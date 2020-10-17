package main

import "fmt"

// Q. Difference between actual and formal parameters

func sum(a []int) int { // parameters received by the function are called Formal Parameters
	total := 0

	for _, v := range a {
		total = total + v
	}
	return total
}

func main() {
	a := []int{40, 60, 10, 2, 5}
	t := sum(a) // parameters passed to function are called Actual Parameter
	fmt.Println(t)
}
