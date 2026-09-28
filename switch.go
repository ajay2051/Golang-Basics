package main

import "fmt"

func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("dive")
	}
}

func main() {
	day := 5
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd days")
	case 2, 4:
		fmt.Println("Even days")
	case 6, 7:
		fmt.Println("Weekends")
	default:
		fmt.Println("Invalid day")
	}
}
