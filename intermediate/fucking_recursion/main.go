package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Println(result)

	sum_of_digit := sumOfDigits(53)
	fmt.Println(sum_of_digit)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

/*

Practical use cases:
- Mathematical algorithms
- Tree and Graphs Traversal
- Divide and Conquer algorithm

Benefits of Recursion:
- Simplicity
- Clarity
- Flexibility

Considerations:
- Performance
- Base case

Best Practices:
- Testing
- Optimization
- Recursive case

*/
