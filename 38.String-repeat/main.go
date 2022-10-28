package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatStr(5, "Hello"))
}
func RepeatStr(repetitions int, value string) string {
	var str string
	for i := 0; i < repetitions; i++ {
		str += value
	}
	return str
}
