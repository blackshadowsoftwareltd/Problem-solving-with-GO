package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetCount("hello world"))
}
func GetCount(str string) (count int) {
	for _, v := range strings.Split(str, "") {
		if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" {
			count++
		}
	}
	return count
}
