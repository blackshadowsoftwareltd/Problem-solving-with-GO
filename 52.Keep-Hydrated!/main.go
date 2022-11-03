package main

import (
	"fmt"
)

func main() {
	fmt.Println(Litres(1234))
}
func Litres(time float64) int {
	return int(time / 2)
}
