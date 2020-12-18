package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	var s square
	s.sideLength = 4

	var t triangle
	t.base = 2
	t.height = 2

	printArea(t)
	printArea(s)
}
