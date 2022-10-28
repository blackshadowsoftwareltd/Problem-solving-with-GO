package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(StringToNumber("1234"))
}
func StringToNumber(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}
