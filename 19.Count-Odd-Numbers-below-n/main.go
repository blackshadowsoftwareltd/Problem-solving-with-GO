package main

import (
	"fmt"
)

func main() {
	fmt.Println(OddCount(15))
}

func OddCount(n int) int {
	return n / 2
}
