package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hero(5, 6))
}
func Hero(bullets, dragons int) bool {
	if bullets >= dragons*2 {
		return true
	} else {
		return false
	}
}
