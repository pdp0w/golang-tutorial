package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error: square root of negative number")
	}

	var ans float64 = x
	for i := 0; i < 1000; i++ {
		result := 0.5 * (ans + x/ans)
		ans = result
	}

	return ans, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}

	return nil
}

func main() {
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(-16))

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	//
	// fmt.Println("Data process successfully")

	// if err := eprocess(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("data read successfully.")
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}

	return nil
}

func readConfig() error {
	return errors.New("Config Error")
}
