package main

import (
	"fmt"
)

func main() {
	fmt.Println(HowMuchILoveYou(6))
}
func HowMuchILoveYou(i int) string {
	for i > 6 {
		i = i - 6
	}
	if i == 1 {
		return "I love you"
	}
	if i == 2 {
		return "a little"
	}
	if i == 3 {
		return "a lot"
	}
	if i == 4 {
		return "passionately"
	}
	if i == 5 {
		return "madly"
	}
	if i == 6 {
		return "not at all"
	}
	return ""
}
