package main

import (
	"fmt"
)

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 8, 7, 4, 5}))
}
func SortNumbers(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	pivot := numbers[0]
	left := []int{}
	right := []int{}
	for _, v := range numbers[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = SortNumbers(left)
	right = SortNumbers(right)
	return append(append(left, pivot), right...)
}
