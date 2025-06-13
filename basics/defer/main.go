package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i:", i)
	defer fmt.Println("Deferred statement executed 1")
	defer fmt.Println("Deferred statement executed 2")
	i++
	fmt.Println("Normal statement executed")
	fmt.Println("Actual value of i:", i)
}

/*
Practical use cases:
- Resource cleanup
- Unlocking mutexes
- Logging and tracing
Best Practices:
- Keep defer action short
- Understand Evalution timing
- Avoid complex control flow
*/
