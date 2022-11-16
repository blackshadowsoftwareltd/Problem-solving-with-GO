package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("4 5 29 54 4 0 -214 542 -64 1 -3 6 -6"))
}
func HighAndLow(in string) string {
	fmt.Println("0" < "1")
	var max, min string
	var m, n, x int
	for _, v := range strings.Split(in, " ") {
		x, _ = strconv.Atoi(v)
		m, _ = strconv.Atoi(max)
		n, _ = strconv.Atoi(min)
		if m == 0 && n == 0 {
			max = v
			min = v
		}
		if x > m {
			max = v
			fmt.Println(v)
		}
		if x < n {
			min = v
		}
	}
	return fmt.Sprintf("%s %s", max, min)
}
