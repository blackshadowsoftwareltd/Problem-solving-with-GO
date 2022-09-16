package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsUpperCase("REMON AHAMMAD"))
}

func IsUpperCase(s string) bool {
	return strings.ToUpper(s) == s
}

///? if it is an interface?
/// ! return strings.ToUpper(string(s)) == string(s)
