package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("abcd", "cd"))
}
func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
