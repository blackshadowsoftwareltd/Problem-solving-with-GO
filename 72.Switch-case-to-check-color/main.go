package main

import "fmt"

func main() {

	var number int
	fmt.Print("Please enter a number : ")
	fmt.Scanln(&number)
	fmt.Printf("You have entered %v", number)
	fmt.Println()
	switch number {
	case 100:
		fmt.Println("Red")
	case 010:
		fmt.Println("Green")
	case 001:
		fmt.Println("Blue")
	case 011:
		fmt.Println("Cyan")
	case 101:
		fmt.Println("Magenta")
	case 110:
		fmt.Println("Yellow")
	case 111:
		fmt.Println("White")
	case 000:
		fmt.Println("Black")
	default:
		fmt.Println("Invalid number")
	}
}
