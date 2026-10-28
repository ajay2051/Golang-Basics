package main

import "fmt"

// Generic functions for any data types string, int etc

type List[T any] struct {
	items []T
}

func Map[T any, U any](l *List[T], fn func(T) U) *List[U] {
	result := List[U]{}
	for _, item := range l.items {
		result.items = append(result.items, fn(item))
	}
	return &result
}

func (l *List[T]) Map[U any](fn func(T) U) *List[U] {
	result := List[U]{}
	for _, item := range l.items {
		result.items = append(result.items, fn(item))
	}
	return &result
}

func add[T int | float64](x T, y T) T {
	return x + y
}

func getValues[K comparable, V any](mp map[K]V) []V {
	var values []V

	for _, v := range mp {
		values = append(values, v)
	}
	return values
}

func printSlices[T any](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}


func main() {

	ages := map[string]int{
		"Ajay":   25,
		"Ram":    30,
		"Shyam":  28,
	}

	ageValues := getValues(ages)
	fmt.Println("Ages:", ageValues)

	// Example 2: map with int keys and string values
	users := map[int]string{
		1: "Ajay",
		2: "Ram",
		3: "Shyam",
	}

	userValues := getValues(users)
	fmt.Println("Users:", userValues)

	// Example 3: map with string keys and float64 values
	prices := map[string]float64{
		"Laptop":  85000.50,
		"Phone":   45000.00,
		"Monitor": 25000.75,
	}

	priceValues := getValues(prices)
	fmt.Println("Prices:", priceValues)

	value := add(2, 3)
	value1 := add(2.5, 3.0)
	fmt.Println(value, value1)

	getValues()

	nums := []int{1, 2, 3}
	names := []string{"golang", "python"}
	printSlices(nums)
}
