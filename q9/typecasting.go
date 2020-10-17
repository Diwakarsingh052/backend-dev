package main

import (
	"fmt"
	"strconv"
)

// Q. Type Cast:
//Int to float, Float to int, Int to string, String to int

func IntToFloat(i int) float64 {
	var f float64
	f = float64(i)
	return f
}
func FloatToInt(f float64) int {
	var i int
	i = int(f)
	return i
}

func IntToString(i int) string {
	var s string

	s = strconv.Itoa(i)

	return s

}

func StringToInt(s string) int {
	var i int
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func main() {
	fmt.Println(IntToFloat(45))
	fmt.Println(FloatToInt(35.9))
	fmt.Println(IntToString(45))
	fmt.Println(StringToInt("32"))
}
