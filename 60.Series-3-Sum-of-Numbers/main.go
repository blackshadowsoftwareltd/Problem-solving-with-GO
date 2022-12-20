package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetSum(0,1))
}
func GetSum(a, b int) int {
	if a == b {
		return a
	}
	sum := 0
	if a < b {
		for i := a; i <= b; i++ {
			sum += i
		}
	} else {
		for i := b; i <= a; i++ {
			sum += i
		}
	}
	return sum
}
