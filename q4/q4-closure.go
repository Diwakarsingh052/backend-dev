package main

//Q- Function closures
import "fmt"

func closure() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	series := closure()

	fmt.Println(series())
	fmt.Println(series())
}
