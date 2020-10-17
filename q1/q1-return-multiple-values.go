package main

import "fmt"

//Q1- . Return multiple values from a function

func multipleValues() (string, int) {
	a, b := 10, 20
	sum := a + b

	return "Total", sum
}

func main() {
	msg, sum := multipleValues()
	fmt.Println(msg, sum)
}
