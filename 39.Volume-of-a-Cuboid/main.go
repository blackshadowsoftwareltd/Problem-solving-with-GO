package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetVolumeOfCuboid(5, 4, 6))
}
func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}
