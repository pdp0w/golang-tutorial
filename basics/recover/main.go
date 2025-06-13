package main

import "fmt"

func main() {
	process()
	fmt.Println("returned from process")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd:", r)
		}
	}()
	fmt.Println("start process")
	panic("something went wrong")
	// fmt.Println("End process")
}

/*

Practical use cases
- Graceful recovery
- clean up
- logging and recovery
Best Practice
- use with defer
- avoid silent recovery
- avoid overuse

*/
