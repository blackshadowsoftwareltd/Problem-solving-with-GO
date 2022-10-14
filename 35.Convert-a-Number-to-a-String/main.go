package main

import (
	"fmt"
)

func main() {
	fmt.Println(NumberToString(5))
}
func NumberToString(n int) string {
	return fmt.Sprintf("%d", n)
}
