package main

import (
	"fmt"
)

func main() {
	fmt.Println(MergeArrays([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}))
}
func MergeArrays(arr1, arr2 []int) []int {
	fmt.Println(arr1)
	fmt.Println(arr2)
	var result []int

	var f bool
	for _, x := range arr1 {
		f = false
		for _, y := range arr2 {
			fmt.Println(x)
			if x == y {
				f = true
				break
			}
		}
		if !f {
			result = append(result, x)
			fmt.Println(x)
		}
	}
	fmt.Println(result)
	return result
}
