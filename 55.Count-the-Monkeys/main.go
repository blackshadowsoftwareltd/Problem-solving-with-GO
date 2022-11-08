package main

import (
	"fmt"
)

func main() {
	fmt.Println(monkeyCount(5))
}
func monkeyCount(n int) []int {
	var monkey []int
	for i := 1; i <= n; i++ {
		monkey = append(monkey, i)
	}
	return monkey
}
