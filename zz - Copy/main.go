package main

import (
	"fmt"
)

var list []any = make([]any, 100)

func main() {
	l := []any{1, 2, []any{3, 4}, 5, 6, []any{7, 8, 9}}
	read(l)
	fmt.Println(list)
}
func read(arr []any) {
	for _, v := range arr {
		if _, ok := v.([]any); ok {
			z := []any{v}
			x := []any{}
			x = append(x, z...)
			read(x)
		}
		list = append(list, v)
	}
}
