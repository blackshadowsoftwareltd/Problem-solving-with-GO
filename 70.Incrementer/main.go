package main

import (
	"fmt"
)

func main() {
	fmt.Println(Incrementer([]int{1, 2, 3, 4, 5}))
}
func Incrementer(n []int) []int {
	for i, v := range n {
		n[i] = (v + i + 1) % 10
	}
	return n
}
