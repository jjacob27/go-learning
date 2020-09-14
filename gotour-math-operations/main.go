package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Math Operations")

	fmt.Println("Sq.root of 2 = ", math.Sqrt(2))
	fmt.Println("Cube root of 2 = ", math.Cbrt(2))
	fmt.Println("e to the power 5 =  ", math.Exp(5))
	fmt.Println("2 to the power 5 = ", math.Exp2(5))

	fmt.Println("Sin 30 degrees = ", math.Sin((30*math.Pi)/180))
}
