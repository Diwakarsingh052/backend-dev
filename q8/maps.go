package main

import "fmt"

//Q. Create a map in Go?

func CreateMap() map[string]string {

	eng := map[string]string{
		"up":   "above",
		"down": "below",
	}
	return eng
}

func main(){
	fmt.Println(CreateMap())
}
