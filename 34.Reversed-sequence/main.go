package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseSeq(5))
}
func ReverseSeq(n int) []int {
	arr := []int{}
	for i := n; 0 < i; i-- {
		arr = append(arr, i)
	}
	return arr
}
