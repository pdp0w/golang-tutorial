package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, base64 encoding")

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", decoded)

	// Url safe, avoid: '/', '+'

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println(urlSafeEncoded)
}

/*

Use cases :
- binary data transfer
- data storage
- embedding resources

Security Consideration :
- It is not encryption
- Proper handeling of padding
- Use appropriate varient

*/
