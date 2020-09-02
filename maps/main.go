package main

import "fmt"

type abc map[string]string

func main() {
	m := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	m["yellow"] = "#abcdef"
	printMap(m)
	fmt.Printf("Length of Map = %d\n", len(m))
	fmt.Println("\n== Checking for existence of white key and then deleting ==")
	if _, contains := m["white"]; contains {
		delete(m, "white")
	}
	printMap(m)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Key = %s, Value =%s\n", key, value)
	}
}
