package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Capitalize("hello world!", []int{1, 4, 6, 12, 13}))
}

func Capitalize(st string, arr []int) string {
	list := strings.Split(st, "")
	for _, v := range arr {
		if v < len(list) {
			list[v] = strings.ToUpper(list[v])
		}
	}
	return strings.Join(list, "")
}
