package main

import "fmt"

func main() {
	// var arrayName [size]elementTypes

	var arr [5]int // go by default initialize elements by zero in array
	fmt.Println(arr)

	arr[0] = 35
	fmt.Println(arr)

	arr[4] = 9
	fmt.Println(arr)

	fruits := [4]string{"apple", "banana", "guava", "graps"}
	fmt.Println("fruits array :", fruits)
	fmt.Println("Third element :", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	//
	// copiedArray[0] = 100
	//
	// fmt.Println("originalArray :", originalArray)
	// fmt.Println("copiedArray :", copiedArray)

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println("element at index,", i, ":", arr[i])
	// }

	for i, v := range arr {
		fmt.Printf("element at index %d : %d\n", i, v) // i = index, v = value
	}

	for _, v := range arr {
		fmt.Printf("element %d\n", v) // _ means discarding the value
	}

	a, _ := someFunction() // _ is discarding b value
	fmt.Println(a)

	fmt.Println("the length of arr array is", len(arr))

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("array1 is equal to array2", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray

	copiedArray[0] = 100

	fmt.Println("originalArray :", originalArray)
	fmt.Println("copiedArray :", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
