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

///? That means. str is 2 D slice.
///? That's why we can use it like str[0][:1]. Represents the first index of value from the first index of the slice.
///? That's why we can use it like str[1][:1]. Represents the first index of value from the second index of the slice.
///? [:1] means the first index of the value.
///? It's split & removes the value from index 1 and returns the value which is before index 1.
