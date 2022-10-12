package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MakeUpperCase("abc"))
}
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
