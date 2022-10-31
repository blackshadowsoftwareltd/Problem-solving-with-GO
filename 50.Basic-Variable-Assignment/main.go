package main

import (
	"fmt"
)

func main() {
	fmt.Println(Namevar())
}
func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}
