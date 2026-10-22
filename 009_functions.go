package main

import "fmt"

func myMessage() {
	fmt.Println("Hello World")
}

func myFamily(fname string, age int) {
	fmt.Println(fname, age)
}

func addition(x int, y int) int {
	return x + y
}

func myFunction(x int, y string) (result int, txt string) {
	result = x + x
	txt = y + "Hello"
	return result, txt
}

// Passing function as a parameter
func getFunc(x string) func(string) string {
	return func(s string) string {
		return x + s
	}

}

func main() {
	myMessage()
	myFamily("Bob", 20)
	value := addition(15, 20)
	fmt.Printf("value: %v\n", value)
	a, b := myFunction(15, "World")
	fmt.Println(a, b)
	f1 := getFunc("ajay")
	fmt.Println(f1("kumar"))
}
