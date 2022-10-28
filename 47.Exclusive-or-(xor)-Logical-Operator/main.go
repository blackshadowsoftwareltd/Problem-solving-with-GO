package main

import (
	"fmt"
)

func main() {
	fmt.Println(Xor(true, true))
}
func Xor(a, b bool) bool {
	return a != b
}
