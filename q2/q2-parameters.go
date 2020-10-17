package main

import "fmt"

// Q2 How many ways can you pass parameters to a method

func passByValue(name string, age int) (string, int) {
	name = "raj"
	age = 21
	return name, age
}

func passByReference(name *string, age *int) {
	*name = "anil"
	*age = 27
}

func main() {
	var (
		name string
		age  int
	)
	name, age = passByValue(name, age)
	fmt.Println(name, age)

	passByReference(&name, &age)
	fmt.Println(name, age)
}
