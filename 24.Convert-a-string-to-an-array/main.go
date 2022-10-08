package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(StringToArray("Hello people how are you?"))
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
