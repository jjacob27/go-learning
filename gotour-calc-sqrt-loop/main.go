package main

import (
	"fmt"
)

func sqrt(n float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - n) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(144))
}
