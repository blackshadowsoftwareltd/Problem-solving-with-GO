package main

import (
	"fmt"
)

func main() {
	fmt.Println(Digitize(0))
}
func Digitize(n int) []int {
	fmt.Println("n", n)
	result := []int{}
	res := 0
	for n > 0 {
		r := n % 10
		result = append(result, (res*10)+r)
		n /= 10
	}
	if len(result) == 0 {
		result = append(result, 0)
	}
	return result
}
