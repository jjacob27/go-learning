package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func prod(x, y int) int {
	return x * y
}

func calculate(x, y int, fn func(int, int) int) int {
	return fn(x, y)
}

func main() {
	myRandSource := rand.NewSource(time.Now().UnixNano())
	myRandGenerator := rand.New(myRandSource)
	myRandomNum := myRandGenerator.Intn(3)
	fmt.Println("Randomly generated number is ", myRandomNum)

	a := 0

	switch myRandomNum {
	case 1:
		fmt.Println("case 1 - add 5,6")
		a = calculate(5, 6, add)
	case 2:
		fmt.Println("case 2 - sub 5,6")
		a = calculate(5, 6, sub)
	case 3:
		fmt.Println("case 3 - prod 5,6")
		a = calculate(5, 6, prod)
	}
	fmt.Println("Result is ", a)

}
