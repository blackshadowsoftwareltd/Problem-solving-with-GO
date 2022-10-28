package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountBy(2, 5))
}

func CountBy(x, n int) []int {
	fmt.Println(x, n)
	var list []int
	for i := 1; i < n+1; i++ {
		list = append(list, x*i)
	}
	return list
}
