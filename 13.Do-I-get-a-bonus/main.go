package main

import (
	"fmt"
)

func main() {
	fmt.Println(BonusTime(100, true))
}
func BonusTime(salary int, bonus bool) string {
	if bonus {
		return fmt.Sprintf("£%d", salary*10)
	} else {
		return fmt.Sprintf("£%d", salary)
	}
}
