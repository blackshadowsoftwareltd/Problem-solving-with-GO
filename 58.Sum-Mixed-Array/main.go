package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumMix([]any{1, 2, 3, 4, 5}))
}
func SumMix(arr []any) int {
	sum := 0
	for _, v := range arr {
		switch v.(type) {
		case int:
			sum += v.(int)
		case string:
			i, _ := strconv.Atoi(v.(string))
			sum += i
		}
	}
	return sum
}
