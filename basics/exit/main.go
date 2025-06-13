package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("starting the main function")

	// exit with the status code 1
	os.Exit(1)

	// this will never be executed
	fmt.Println("end of main function")
}

/*
Practical use cases :
- Error handeling
- termination condition
- exit codes
Best practice :
- avoid deferred actions
- status code
- avoid abusive use
*/
