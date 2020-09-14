package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My random number is ", rand.Intn(1000))

	fmt.Println("Doing random numbers with a different fixed seed ")
	myRandSource := rand.NewSource(50)
	myRand := rand.New(myRandSource)
	fmt.Println("Second Rand Number is ", myRand.Intn(1000))

	myRandSource = rand.NewSource(time.Now().UnixNano())
	myRand = rand.New(myRandSource)
	fmt.Println("Third Rand Number is ", myRand.Intn(1000))
}
