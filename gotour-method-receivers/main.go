package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v Vertex) scale(scaleFactor float64) {
	v.x *= scaleFactor
	v.y *= scaleFactor
}

func (v *Vertex) scalePointers(scaleFactor float64) {
	v.x *= scaleFactor
	v.y *= scaleFactor
}

func main() {
	v1 := Vertex{3, 4}
	v1.scale(3)
	fmt.Println(v1)

	v1.scalePointers(3)
	fmt.Println(v1)
}
