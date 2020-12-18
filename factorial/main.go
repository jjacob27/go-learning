package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number :")

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading input")
	}

	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	number, err1 := strconv.Atoi(str)

	if err1 != nil {
		fmt.Println("Error in converting to number")
	}

	if number < 0 {
		fmt.Println("Factorial of negative nos is not supported")
		return
	}

	if number > 20 {
		fmt.Println("Can't calculate factorials of numbers beyond 20")
		return
	}

	var fact int64
	fact = factorial(int64(number))
	fmt.Println("Factorial is ", fact)
}
