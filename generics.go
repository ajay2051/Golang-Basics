package main

import "fmt"

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

func main() {
	value := add(2, 3)
	value1 := add(2.5, 3.0)
	fmt.Println(value, value1)
}
