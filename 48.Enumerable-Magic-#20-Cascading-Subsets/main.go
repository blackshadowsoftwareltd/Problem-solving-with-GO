package main

import (
	"fmt"
)

func main() {
	fmt.Println(EachCons([]int{1, 2, 3, 4, 5, 6}, 3))
}
func EachCons(arr []int, n int) [][]int {
	var result [][]int
	for i := n; i <= len(arr); i++ {
		result = append(result, arr[i-n:i])
	}
	return result
}
