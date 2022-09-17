package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AbbrevName("REMON ahammad"))
}
func AbbrevName(name string) string {
	str := strings.Split(name, " ")
	return strings.ToUpper(str[0][:1]) + "." + strings.ToUpper(str[1][:1])
}
