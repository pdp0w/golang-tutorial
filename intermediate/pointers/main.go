package main

import "fmt"

func main() {

	// var ptr *int

	var ptr *int
	var a int = 10

	ptr = &a

	fmt.Println(ptr)
	fmt.Println(*ptr) // dereference
	// zero values of pointer is nil pointer

	// if ptr == nil {
	// 	fmt.Println("pointer is nil")
	// }

	modifyValue(ptr)
	fmt.Println(*ptr, a)
}

func modifyValue(ptr *int) {
	*ptr++
}
