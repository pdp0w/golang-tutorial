package main

import "fmt"

func main() {
	// sequence := adder()
	//
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	//
	// sequence2 := adder()
	//
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())

	subtraction := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtraction(1))
	fmt.Println(subtraction(1))
	fmt.Println(subtraction(1))
	fmt.Println(subtraction(1))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

/*

Pratical use cases:
- Stateful function
- Encapsulation
- Callbacks

Usefulness of closures:
- Encapsulation
- Flexibility
- Readability

Considerations:
- Memory usage
- Concurrency

Best Practice:
- Limit scope
- Avoid Overuse

*/
