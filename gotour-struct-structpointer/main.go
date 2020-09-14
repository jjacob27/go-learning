package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("Original v=", v)
	v.x = 5
	fmt.Println("After v.x change=", v)

	p := &v
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	(*p).x = 6
	fmt.Println("After pointer change *p = ", *p)
	fmt.Println("After pointer change=", v)
}
