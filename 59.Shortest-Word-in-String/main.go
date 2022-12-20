package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("Hello, World! What's up?"))
}
func FindShort(s string) int {
	data := strings.Split(s, " ")
	x := len(s)
	for _, v := range data {
		if x > len(v) || x == len(v) {
			x = len(v)
		}
	}
	return x
}
