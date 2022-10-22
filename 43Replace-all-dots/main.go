package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReplaceDots("djjfal.d adf ad.fj"))
}
func ReplaceDots(str string) string {
	return strings.ReplaceAll(str, ".", "-")
}
