package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("assa"))
}
func IsPalindrome(str string) bool {
	var flag bool = true
	arr := strings.Split(str, "")
	for i, v := range arr {
		if strings.ToLower(v) != strings.ToLower(arr[len(arr)-i-1]) {
			flag = false
			break
		}
	}
	return flag
}
