package main

import "fmt"

func main() {
	i := 50
	var p *int

	p = &i
	fmt.Println("i=", i)
	fmt.Println("p=", p)
	fmt.Println("&i=", &i)
	fmt.Println("*p=", *p)
}
