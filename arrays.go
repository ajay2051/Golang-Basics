package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [2]string{"hello", "world"}
	fmt.Println(arr1, arr2)
}

// Inferred Lengths

func main() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]string{"hello", "world"}
	fmt.Println(arr1, arr2)
}

// Access elements of array
func main() {
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

// Change elements of array
func main() {
	var arr1 = [...]int{1, 2, 3}
	arr1[1] = 10
	fmt.Println(arr1)
}

// Initialize specific elements
func main() {
	arr1 := [5]int{1: 10, 2: 20, 3: 30}
	fmt.Println(arr1)
}

// Length of an array
func main() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
