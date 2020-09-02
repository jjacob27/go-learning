package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Print("Enter a number n (0 - 10): ")
	_, error := fmt.Scanf("%d", &number)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	for i := 0; i <= number; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
