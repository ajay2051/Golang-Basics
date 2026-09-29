package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// looping through strings
func main() {
	str := "hello world"

	fmt.Println(string(str[0]))
}

func main() {
	str := "hello world"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	for _, char := range str {
		fmt.Printf("%c", char)
		break
	}
}
