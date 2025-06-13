package main

import "fmt"

func main() {

	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType {}

	// fmt.Println(sum(2, 3, 4, 5, 6, 7))

	// statement, total := sum("The sum of 1 2 3 is", 1, 2, 3)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	statement, total := sum("sum of numbers", numbers...)

	fmt.Println(statement, total)
}

func sum(returnString string, nums ...int) (string, int) { // variadic parameters should come at last
	total := 0
	for _, v := range nums {
		total += v
	}

	return returnString, total
}
