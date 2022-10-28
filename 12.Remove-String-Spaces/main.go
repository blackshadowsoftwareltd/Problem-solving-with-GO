package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NoSpace("Remon Ahammad"))
}
func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
