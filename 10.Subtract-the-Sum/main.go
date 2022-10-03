package main

import (
	"fmt"
)

func main() {
	fmt.Println(SubtractSum(9599))
}

func SubtractSum(n int) string {
	var m, x, sum int
	x = n
	for n > 100 {
		for x > 0 {
			m = x % 10
			sum = sum + m
			x = x / 10
		}
		n = n - sum
	}
	fmt.Println(n)
	arr := []string{
		"kiwi",
		"pear",
		"kiwi",
		"banana",
		"melon",
		"banana",
		"melon",
		"pineapple",
		"apple",
		"pineapple",
		"cucumber",
		"pineapple",
		"cucumber",
		"orange",
		"grape",
		"orange",
		"grape",
		"apple",
		"grape",
		"cherry",
		"pear",
		"cherry",
		"pear",
		"kiwi",
		"banana",
		"kiwi",
		"apple",
		"melon",
		"banana",
		"melon",
		"pineapple",
		"melon",
		"pineapple",
		"cucumber",
		"orange",
		"apple",
		"orange",
		"grape",
		"orange",
		"grape",
		"cherry",
		"pear",
		"cherry",
		"pear",
		"apple",
		"pear",
		"kiwi",
		"banana",
		"kiwi",
		"banana",
		"melon",
		"pineapple",
		"melon",
		"apple",
		"cucumber",
		"pineapple",
		"cucumber",
		"orange",
		"cucumber",
		"orange",
		"grape",
		"cherry",
		"apple",
		"cherry",
		"pear",
		"cherry",
		"pear",
		"kiwi",
		"pear",
		"kiwi",
		"banana",
		"apple",
		"banana",
		"melon",
		"pineapple",
		"melon",
		"pineapple",
		"cucumber",
		"pineapple",
		"cucumber",
		"apple",
		"grape",
		"orange",
		"grape",
		"cherry",
		"grape",
		"cherry",
		"pear",
		"cherry",
		"apple",
		"kiwi",
		"banana",
		"kiwi",
		"banana",
		"melon",
		"banana",
		"melon",
		"pineapple",
		"applpineapple",
	}
	return fmt.Sprintf("%d.%s", n, arr[n])
}
