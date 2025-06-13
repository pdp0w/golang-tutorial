package main

import (
	"errors"
	"fmt"
)

func main() {
	// function funcName(parameter1, parameter2, ...) (returnType1, returnType2, ...) {}

	q, r := divide(10, 3)
	fmt.Println(q, r)

	q1, r1 := divide2(10, 3)
	fmt.Println(q1, r1)

	result, err := compare(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func divide2(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
