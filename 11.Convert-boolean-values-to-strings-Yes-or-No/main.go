package main

import (
	"fmt"
)

func main() {
	fmt.Println(BoolToWord(true))
}
func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	} else {
		return "No"
	}
}
