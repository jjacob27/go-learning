package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing file arguement")
		os.Exit(1)
	}

	filename := os.Args[1]
	filename = strings.TrimSpace(filename)

	fmt.Println("Attempting to open file = ", filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open file ")
		os.Exit(1)
	}

	defer f.Close()
	fmt.Println("Contents of file are ")
	io.Copy(os.Stdout, f)
}
