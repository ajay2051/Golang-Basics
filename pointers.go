package main

import "fmt"

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func change(x *int) {
	*x = 100
}

func main() {
	x := 0
	y := &x

	*y = 100

	a := 10
	change(&a)
	fmt.Println(a)
	fmt.Println(x, *y)

	b := Book{id: 1, title: "book1"}
	b.setTitle("book2")
	fmt.Println(b)
}
