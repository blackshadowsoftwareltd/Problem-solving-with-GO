package main

import (
	"fmt"
)

func main() {
	fmt.Println(Collatz(20))
}
func Collatz(n int) int {
    sequence := []int{n}
    for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n*3 + 1
        }
        sequence = append(sequence, n)
    }
    return len(sequence)
}