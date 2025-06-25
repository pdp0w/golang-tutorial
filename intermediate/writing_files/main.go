package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file...")
		return
	}

	defer file.Close()

	// write data to file
	data := []byte("Hello, World!\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing file")
		return
	}

	fmt.Println("data has been writen to file.")

	file, err = os.Create("writingString.txt")
	if err != nil {
		fmt.Println("error creating file.")
		return
	}

	defer file.Close()

	_, err = file.WriteString("Hello, Go\n")
	if err != nil {
		fmt.Println("Error writing file.")
		return
	}

}
