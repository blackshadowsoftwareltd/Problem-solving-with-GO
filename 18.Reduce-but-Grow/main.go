package main

import (
	"fmt"
)

func main() {

	fmt.Println(Grow([]int{1, 2, 3, 4, 5}))
}
func Grow(arr []int) int {
	result := 1
	for i := 0; i < len(arr); i++ {
		result = result * arr[i]
	}
	return result
}
