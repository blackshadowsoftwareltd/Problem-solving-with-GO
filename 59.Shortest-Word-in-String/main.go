package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("Hello, World! What's up?"))
}
func FindShort(s string) int {
	x := len(s)
	for _, v := range strings.Split(s, " ") {
		if x > len(v) {
			x = len(v)
		}
	}
	return x
}
