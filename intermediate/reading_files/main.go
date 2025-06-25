package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../writing_files/output.txt")
	if err != nil {
		fmt.Println("error opening file.")
		return
	}

	defer func() {
		fmt.Println("Closing file...")
		file.Close()
		fmt.Println("Files closed...")
	}()

	fmt.Println("file opened successfully.")

	// read the contents of the files
	// data := make([]byte, 1024)
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading file...")
	// 	return
	// }
	//
	// fmt.Println(string(data))

	scanner := bufio.NewScanner(file)

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line:", line)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("error reading file...")
		return
	}

}
