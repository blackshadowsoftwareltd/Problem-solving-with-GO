package main

import (
	"fmt"
)

func main() {
	fmt.Println(Smaller([]int{5, 4, 3, 2, 1}))
}

func Smaller(arr []int) []int {
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		count := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				count++
			}
		}
		result[i] = count
	}
	return result
}
