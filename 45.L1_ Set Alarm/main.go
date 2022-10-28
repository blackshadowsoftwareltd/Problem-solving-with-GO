package main

import (
	"fmt"
)

func main() {
	fmt.Println(SetAlarm(true, false))
}
func SetAlarm(employed, vacation bool) bool {
	if employed == true && vacation == false {
		return true
	}
	return false
}
