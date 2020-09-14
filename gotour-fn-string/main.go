package main

import (
	"fmt"
)

func reverse(input string) string {
	n := 0
	runes := make([]rune, len(input))
	for _, r := range input {
		runes[n] = r
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println(reverse(input))
}
