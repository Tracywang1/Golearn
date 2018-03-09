package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

func main() {
	fmt.Println("hello")
	var a Integer = 1
	var b LessAdder = &a
	// var c LessAdder = a
	fmt.Println(a, b)
	var c Integer = 1
	var d1 Lesser = &c
	var d2 Lesser = c
	fmt.Println(d1, d2)
}
