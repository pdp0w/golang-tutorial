package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// conditional statement
	if 5 == 5 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	} // same as c

	/*
		if (condition) {
			body...
		} else if (condition) {
			body...
		} else {
			body...
		}
	*/

	// switch cases condition things

	/*
		switch expression {
		case value1:
			// code
			fallthrough // this does not break and hence continue
		case value2:
			// code
		default:
			// code
		}

		here is the catch, unlike other language go break automatically when case statisfied, If one want to continue executing below then use "fallthrough"
	*/

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("Booyaah")
		fallthrough
	case "banana":
		fmt.Println("Nooyaah")
	default:
		fmt.Println("Mmmyaah")
	}

	day := "sunday"

	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Weekdays")
	case "sunday", "saturday":
		fmt.Println("Weekends")
	default:
		fmt.Println("invalid day")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is greater than 20")
	}

	checkType(40)
	checkType(40.09343)
	checkType("pdp0w")
}

func checkType(x interface{}) { // x interface{} meaning it can take any data types
	switch x.(type) {
	case int:
		fmt.Println("it is int")
	case float64:
		fmt.Println("it is float64")
	case string:
		fmt.Println("it is string")
	default:
		fmt.Println("unknown")
	}
}
