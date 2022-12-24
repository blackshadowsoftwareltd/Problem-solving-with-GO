package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solve("[]any{1, 2, 3, 4, 5}"))
}
func Solve(s string) []int {
	var uppercase, lowercase, numbers, special int
	for _, v := range s {
		switch {
		case v >= 'A' && v <= 'Z':
			uppercase++
		case v >= 'a' && v <= 'z':
			lowercase++
		case v >= '0' && v <= '9':
			numbers++
		case strings.ContainsAny(string(v), "!@#$%^&*()_+-=[]{}|\\:;'\"<>,./?"):
			special++
		}
	}
	return []int{uppercase, lowercase, numbers, special}
}
