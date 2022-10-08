package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BinToDec("1000101010101010101010010010001001"))
}

func BinToDec(bin string) int {
	dec, err := strconv.ParseUint(bin, 2, 64)
	if err != nil {
		panic(err)
	} 
	return int(dec)
}
