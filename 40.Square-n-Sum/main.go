package main

import (
	"fmt"
)

func main() {
	fmt.Println(SquareSum([]int{1, 6, 4, 7, 2}))
}
func SquareSum(numbers []int) int {
	var result int
	for _, v := range numbers {
		result += v * v
	}
	return result
}
