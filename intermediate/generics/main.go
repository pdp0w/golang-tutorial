package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("The stack is empty")
		return
	}

	fmt.Print("Stack Elements : ")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	// a := 3
	// b := 2
	//
	// fmt.Println(a, b)
	//
	// swap(a, b)
	//
	// fmt.Println(a, b)

	intStack := Stack[int]{}

	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printAll()

	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())

	intStack.printAll()

	stringStack := Stack[string]{}

	stringStack.push("pdp0w")
	stringStack.push("Ranveer Kapoor")
	stringStack.push("SRK")

	stringStack.printAll()

}

/*

Generics :
- benifits of generics :
	- Code reusability
	- Type safety
	- Performance
- Consideration :
	- Type consideration
	- Documentation
	- Testing

*/
