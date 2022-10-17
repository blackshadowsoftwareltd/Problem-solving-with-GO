package main

import (
	"fmt"
)

func main() {
	fmt.Println(MergeArrays([]int{9, 1, 2, 5}, []int{1, 6, 4, 7, 2}))
}

func MergeArrays(arr1, arr2 []int) []int {
	list := arr2
	var f bool
	for _, x := range arr1 {
		f = false
		for _, y := range arr2 {
				if x == y {
				f = true
				break
			}
		}
		if !f {
			list = append(list, x)
		}
	}
	return list
}
