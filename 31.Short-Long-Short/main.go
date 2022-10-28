package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("1", "22"))
}

func Solution(a, b string) string {
	if len(a) < len(b) {
		return fmt.Sprintf("%s%s%s", a, b, a)
	}
	return fmt.Sprintf("%s%s%s", b, a, b)
}
