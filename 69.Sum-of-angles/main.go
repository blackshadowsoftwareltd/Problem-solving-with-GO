package main

import (
	"fmt"
)

func main() {
	fmt.Println(Angle(5))
}
func Angle(n int) int {
    return (n-2)*180
}
