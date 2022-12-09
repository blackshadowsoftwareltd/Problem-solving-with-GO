package main

import (
	"fmt"
)

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))
}
func PositiveSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
