package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strSlices := strings.Split(s," ")
	
	var m = map[string]int{}
	
	for _,v := range strSlices {
		e := m[v];
		e = e+1
		m[v] = e
	}
	
	return m
}

func main() {
	fmt.Println(WordCount("I am feeling super and I am feeling elated and nice, so do I"))
}
