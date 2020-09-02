package main

import "fmt"

type person struct {
	name           string
	age            int
	contactDetails contact
}

type contact struct {
	email string
	phone string
}

func main() {
	s := "abc"
	fmt.Printf("The string printed as is %s\n", s)
	fmt.Printf("Address of string printed =%s\n", &s)
	print(s)
	printPointer(&s)

}

func printPointer(sCopy *string) {
	fmt.Printf("String pased by pointer, value is %s\n", *sCopy)
	fmt.Printf("String pased by pointer, address is %s\n", &sCopy)
}

func print(sCopy string) {
	fmt.Printf("String passed by value to function, value is %s\n", sCopy)
	fmt.Printf("String passed by value to function, address is %s\n", &sCopy)
}
