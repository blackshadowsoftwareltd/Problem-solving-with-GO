package main

import (
	"fmt"
)

func main() {
	fmt.Println(Summation(8))
}

func Summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}
