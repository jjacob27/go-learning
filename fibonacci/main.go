package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Print("Enter a number n for generating fibonacci: ")
	_, error := fmt.Scanf("%d", &number)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	f := fibo(number)

	fmt.Println(f)

}

func fibo(n int) []int {
	f := make([]int, n+1, n+2)

	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
