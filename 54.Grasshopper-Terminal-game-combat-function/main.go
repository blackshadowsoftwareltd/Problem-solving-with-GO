package main

import (
	"fmt" 
)

func main() {
	fmt.Println(combat(1,2))
}
func combat(health, damage float64) float64 {
  if health-damage>0{
    return health-damage
  }
  return 0
}
