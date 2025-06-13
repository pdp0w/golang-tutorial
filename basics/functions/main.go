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

	result := applyOperation(5, 3, add)
	fmt.Println(result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println(multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}
