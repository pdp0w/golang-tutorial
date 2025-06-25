package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	joinPath := filepath.Join("downloads", "file.zip")
	fmt.Println(joinPath)

	dir, file := filepath.Split("/home/priyanshoon/avnish.pdf")
	fmt.Println(file)
	fmt.Println(dir)
	fmt.Println(filepath.Base("/home/priyanshoon/avnish.pdf"))

}
