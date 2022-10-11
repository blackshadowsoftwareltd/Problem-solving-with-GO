package main

import (
	"fmt"
)

func main() {
	fmt.Println(century(1990))
}
func century(year int) int {
	if year%100 > 0 {
		return (year / 100) + 1
	}
	return year / 100
}
