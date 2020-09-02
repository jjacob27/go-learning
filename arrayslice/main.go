package main

import "fmt"

func main() {
	strslice := []string{"a", "b", "c", "d", "e"}

	fmt.Println(len(strslice))
	strslice = append(strslice, "g")
	fmt.Println(len(strslice))

	for i, str := range strslice {
		fmt.Println(str)
	}
}
