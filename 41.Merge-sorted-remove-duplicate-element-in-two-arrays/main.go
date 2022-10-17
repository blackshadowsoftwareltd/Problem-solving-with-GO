package main

import (
	"fmt"
)

func main() {
	fmt.Println(MergeArrays([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))
}

func MergeArrays(arr1, arr2 []int) []int {

	fmt.Println(arr1)
	fmt.Println(arr2)

	var i, temp int
	var flag bool
	///? Merge the two arrays
	for _, x := range arr1 {
		flag = false
		for _, y := range arr2 {
			if x == y {
				flag = true
				break
			}
		}
		if !flag {
			arr2 = append(arr2, x)
		}
	}
	arr1 = []int{}
	///? Sorting
	keys := make(map[int]bool)
	for j := 0; j < len(arr2); j++ {
		for i = 0; i < len(arr2); i++ {
			if i == len(arr2)-1 {
				break
			}
			if arr2[i] > arr2[i+1] {
				temp = arr2[i]
				arr2[i] = arr2[i+1]
				arr2[i+1] = temp
			}
		}
	}
	///? remove duplicate
	for _, v := range arr2 {
		if _, value := keys[v]; !value {
			keys[v] = true
			arr1 = append(arr1, v)
		}
	}
	return arr1
}
