package main

import (
	"fmt"
)

func main() {
	fmt.Println(SmallestIntegerFinder([]int{1, 2, -8, 4, 5}))
}
func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if smallest > numbers[i] {
			smallest = numbers[i]
		}
	}
	return smallest
}
