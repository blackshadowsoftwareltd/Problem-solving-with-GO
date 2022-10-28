package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("My name is Remon Ahammad"))
}
func ReverseWords(str string) string {
	words := strings.Fields(str)
	var resutl []string
	for i := len(words) - 1; i >= 0; i-- {
		resutl = append(resutl, words[i])
	}
	return strings.Join(resutl, " ")

}
