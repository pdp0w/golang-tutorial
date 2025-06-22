package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := doSomething(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Operation completed successfully")
}

type customError struct {
	code    int
	message string
	err     error
}

// Error returns the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.err)
}

// function that returns a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong!",
// 	}
// }

func doSomething() error {
	if err := doSomethingElse(); err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}

	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}
