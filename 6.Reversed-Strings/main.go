package main

import (
	"strings"
)

func main() {
	fmt.Println(Solution("REMON"))
}
func Solution(word string) string {
	var result []string
	str := strings.Split(word, "")
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return strings.Join(result, "")
}
