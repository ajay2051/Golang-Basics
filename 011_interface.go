package main

import "fmt"

type Shape interface {
	getPerimeter() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

type Square struct {
	length  float64
	breadth float64
}

func (t Triangle) getPerimeter() float64 {
	return t.a + t.a + t.b
}

func (s Square) getPerimeter() float64 {
	return s.length * s.length * s.breadth
}

func main() {
	var s Shape = Triangle{1.0, 2.0, 3.0}
	var a Shape = Square{3.0, 4.0}
	fmt.Println(s.getPerimeter())
	fmt.Println(a.getPerimeter())
}
