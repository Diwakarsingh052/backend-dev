package main

import "fmt"

//Q- Variable declaration and variable definition

func variables() (int, string) {
	var i int // declaring a variable and by default initialized with 0
	i = 10    // definition
	_ = i

	// declaration of variables and intialized with default values
	var (
		name string // default value empty string ""
		age  int    // default value 0
	)
	_ = name
	_ = age
	return age, name

}

func main() {
	fmt.Println(variables())
}
