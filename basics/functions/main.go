package main

import "fmt"

func main() {
	/*
		func <func_name>(parameter_list) (output returnType) {
			statement...
		}

		if public function -> start with uppercase letter
		if private function -> start with lowercase letter
	*/

	fmt.Println(add(2, 2))

	func() {
		fmt.Println("hello from anonymous function")
	}()

	greet := func() {
		fmt.Println("hello from anonymous function")
	}

	greet()
}

func add(a, b int) int {
	return a + b
}
