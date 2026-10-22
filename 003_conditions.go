package main

import "fmt"

func main() {
	x := 2
	if x < 5 {
		fmt.Println("x is smaller than 5")
	} else if x > 1 {
		fmt.Println("x is lower than 10")
	} else {
		fmt.Println("x is greater than 10")
	}

}

func main() {
	num := 20
	if num > 10 {
		fmt.Println("num is smaller than 10")
		if num > 15 {
			fmt.Println("num is lower than 15")
		}
	} else {
		fmt.Println("num is greater than 10")
	}

}
