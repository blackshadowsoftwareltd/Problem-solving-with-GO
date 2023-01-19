package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(HasUniqueChar("  Assa"))
}
func HasUniqueChar(str string) bool {
	fmt.Println(str)
	s := strings.Split(str, "")
	f := false
	for i, v := range s {

		for j, x := range s {

			if i == j {
				continue
			} else if v == x {
				f = true
				break
			}
		}
	}
	return !f
}
