package main

import "fmt"

func main() {
	// panic(interface{})

	// Example of valid input
	process(10)

	// Example of invalid input
	process(-10)
}

func process(input int) {

	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")

	if input < 0 {
		fmt.Println("before panic")
		panic("input must be a non-negative number")
	}

	fmt.Println("Processing input:", input)
}
