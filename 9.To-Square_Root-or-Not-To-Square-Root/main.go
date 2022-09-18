package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{4, 3, 9, 7, 2, 1}
	fmt.Println(SquareOrSquareRoot(slice))
}

func SquareOrSquareRoot(arr []int) []int {
	result := []int{}
	for _, v := range arr {
		if math.Sqrt(float64(v)) == float64(int(math.Sqrt(float64(v)))%v) {
			result = append(result, int(math.Sqrt(float64(v))))
		} else {
			result = append(result, v*v)
		}
	}
	return result
}
