package main

import (
	"fmt"
)

func main() {

	fmt.Println(OddCount(7))
}

func OddCount(n int) int {
	var arr []int

	for i := n - 1; i > 0; i-- {
		if i%2 != 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
	for len(arr) != 1 {
		arr = arr[:1]
		fmt.Println(arr)
		if len(arr) != 1 {
			arr = arr[len(arr)-1:]
			fmt.Println(arr)
		}
	}
	return arr[0]
}
