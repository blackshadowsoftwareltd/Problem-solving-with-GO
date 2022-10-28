package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetGrade(40, 60, 80))
}
func GetGrade(a, b, c int) rune {
	var x = (a + b + c) / 3
	if x >= 90 {
		return 'A'
	} else if x >= 80 {
		return 'B'
	} else if x >= 70 {
		return 'C'
	} else if x >= 60 {
		return 'D'
	} else {
		return 'F'
	}
}
