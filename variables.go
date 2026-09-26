package main

import "fmt"

func main() {
	var a string = "Hello"
	var b string = "World"

	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	var name string
	name = "Jack"

	fmt.Println(name)
}

func main() {

	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c)
}

func main() {
	var a, b = 1, "hello"
	c, d := 2, "world"
	fmt.Println(a, b, c, d)
}
