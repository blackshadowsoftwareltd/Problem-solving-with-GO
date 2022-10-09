package main

import (
	"fmt"
)

func main() {
	fmt.Println(Between(8, 20))
}
func Between(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}
