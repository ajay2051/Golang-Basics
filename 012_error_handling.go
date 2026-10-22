package main

import (
	"errors"
	"fmt"
)

//func divide(a int, b int) int {
//	return a / b
//}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func defferedFunc() {
	fmt.Println("defferedFunc")
	r := recover()
	if r == nil {
		fmt.Println(r)
	} else {
		fmt.Println(r)
	}
}

func main() {
	defer defferedFunc()
	result, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	panic("this should never happen")
}
