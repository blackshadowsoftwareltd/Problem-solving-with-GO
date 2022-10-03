package main

import (
	"fmt"
)

func main() {
	fmt.Println(QuarterOf(6))
}
func QuarterOf(month int) int {
	  months  := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	return months[month-1]
}
