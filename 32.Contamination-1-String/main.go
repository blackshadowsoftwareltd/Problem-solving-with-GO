package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Contamination("abc", "z"))
}
func Contamination(text, char string) string {
	arr := []string{}
	for i := 0; i < len(text); i++ {
		arr = append(arr, char)
	}
	return strings.Join(arr, "")
}
