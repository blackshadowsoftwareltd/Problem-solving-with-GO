package main

import (
	"fmt"
)

func main() {
	fmt.Println(Invert([]int{1, 2, -3, 4, 5}))
}
func Invert(arr []int) (result []int) {
	for _, v := range arr {
		result = append(result, v*-1)
	}
	return result
}
