package main

import "fmt"

func main() {
	var s Stack
	s.Push("apple")
	s.Push("mango")
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.IsEmpty())
	s.Pop()
	fmt.Println(s.IsEmpty())
}
