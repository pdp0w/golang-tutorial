package main

import "fmt"

func init() {
	fmt.Println("initializing package...")
}

func init() {
	fmt.Println("initializing second package...")
}

func main() {
	fmt.Println("Inside the main function")
}

/*

practical use cases :
- setup task
- configurations
- registering components
- database initialization
Best practice
- avoid side effects
- initialization order
- documentation

*/
