package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToJadenCase("Hello World"))
}
func ToJadenCase(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}
