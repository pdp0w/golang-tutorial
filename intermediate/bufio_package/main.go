package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("Hello, bufio package yo yo yo!\n"))
	//
	// // Reading byte slice
	// data := make([]byte, 5)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Printf("Read %d bytes: %s\n", n, data)
	//
	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println(line)

	writer := bufio.NewWriter(os.Stdout)

	// writing byte slice

	data := []byte("Hello, bufio package!\n")
	nn, err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(nn)

	// flush the buffer to ensure all data is written to os.Stdout

	if err = writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	str := "This is the string\n"
	n, err := writer.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)

	if err = writer.Flush(); err != nil {
		fmt.Println(err)
		return
	}

}

/*

Use cases and benifits
- buffering
- convenience methods
- error handeling

Best practices
- Check errors
- wrap reader and writer instance for effiecient buffered I/O operations
- don't forget to call flush

*/
