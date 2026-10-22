package main

import "fmt"

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

func main() {
	value := add(2, 3)
	value1 := add(2.5, 3.0)
	fmt.Println(value, value1)
}
