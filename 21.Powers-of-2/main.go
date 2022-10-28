package main

import (
	"fmt"
)

func main() {
	fmt.Println(PowersOfTwo(10))
}
func PowersOfTwo(n int) []uint64 {
	var list []uint64

	for i := 0; i <= n; i++ {
		list = append(list, uint64(1<<i))
	}
	return list
}
