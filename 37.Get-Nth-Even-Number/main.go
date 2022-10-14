package main

import (
	"fmt"
)

func main() {
	fmt.Println(NthEven(5))
}
func NthEven(n int) int {
	return (n * 2) - 2
}
