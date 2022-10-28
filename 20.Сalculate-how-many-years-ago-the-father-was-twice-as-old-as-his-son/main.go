package main

import (
	"fmt"
)

func main() {
	fmt.Println(TwiceAsOld(55, 30))
}
func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	sonYearsOld = sonYearsOld * 2
	temp := dadYearsOld - sonYearsOld
	if temp > 0 {
		return temp
	} else {
		return temp + temp*-2
	}
}
