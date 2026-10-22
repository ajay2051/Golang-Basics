package main

import (
	"fmt"
	"math"
)

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}

func main() {
	x := 2
	y := 3
	z := x + y
	fmt.Println(z)
}

func main() {
	fmt.Println(math.Min(3, 5))
}
