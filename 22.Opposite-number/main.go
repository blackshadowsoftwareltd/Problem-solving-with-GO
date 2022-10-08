package main

import (
	"fmt"
)

func main() {
	fmt.Println(Opposite(-10))
}

func Opposite(value int) int {
	if value > 0 {
		return -value
	} else {
		return value * -1
	}
}
