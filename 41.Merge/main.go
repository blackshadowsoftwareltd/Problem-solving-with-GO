package main

import (
	"fmt"
)

func main() {
	fmt.Println(MergeArrays([]int{-98, -91, -82, -79, -77, -74, -64, -61, -61, -58, -56, -49, -48, -48, -45, -36, -34, -24, 2, 10, 11, 11, 16, 18, 18, 27, 29, 50, 67, 74, 78, 91, 92, 95}, []int{-99, -98, -95, -93, -84, -82, -82, -80, -75, -74, -74, -73, -70, -69, -66, -65, -56, -52, -51, -46, -45, -43, -41, -39, -36, -36, -35, -34, -33, -32, -30, -29, -26, -25, -24, -21, -20, -18, -17, -11, -9, -8, -6, 2, 5, 5, 7, 8, 12, 13, 13, 13, 17, 17, 19, 19, 19, 24, 24, 25, 27, 29, 30, 34, 36, 36, 38, 38, 40, 41, 41, 42, 42, 43, 47, 48, 50, 51, 53, 57, 74, 82, 84, 87, 87, 89, 94, 94, 100}))
}

func MergeArrays(arr1, arr2 []int) []int {

	fmt.Println(arr1)
	fmt.Println(arr2)

	var i, temp int
	var flag bool
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
	return arr2
}
