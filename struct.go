package main

import "fmt"

type Sport struct {
	name     string
	position string
}

type Person struct {
	name     string
	age      int
	favSport Sport
	sports   []Sport
}

// name can be imported, but Name cannot be imported. variables starting with Capital letter can't be imported.

func (p Person) getName() string {
	return p.name
}

func (p Person) setName(s string) {
	p.name = s
}

func main() {
	p1 := Person{"tim", 23, Sport{"soccer", "defender"}, []Sport{Sport{"cricket", "batter"}}}
	p1.name = "jack"
	value := p1.getName()
	fmt.Println(p1)
	fmt.Println(value)
	p1.setName("jay")
	fmt.Println(p1)
	fmt.Println(p1.favSport)
}
