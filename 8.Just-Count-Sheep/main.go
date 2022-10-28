package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSheep(3))
}

///? first way
func countSheep(num int) string {
	var str string
	for i := 0; i < num; i++ {
		str += fmt.Sprintf("%d sheep...", i+1)
	}
	return str
}

///? second way
func countSheep(num int) string {
	var str []string
	for i := 0; i < num; i++ {
		str = append(str, fmt.Sprintf("%d sheep...", i+1))
	}
	return strings.Join(str, "")
}
