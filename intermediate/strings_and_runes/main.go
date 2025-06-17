package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, World!"
	rawMessage := `Hello\nGo` // in backslash, escape sequence will not work, like \
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(rawMessage)
	fmt.Println(message2)

	// strings are array of unicode char

	fmt.Println("length of message variable is", len(message))

	fmt.Println("The first character is message var is", message[0])

	greetings := "Hello "
	name := "Rounak"

	msg := greetings + name

	fmt.Println(msg)

	str1 := "Apple"
	str2 := "Banana"
	fmt.Println(str1 < str2)

	for _, chat := range message {
		fmt.Printf("%x\n", chat) // print hexa decimal value
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greetings))

	greetingsWithName := greetings + name
	fmt.Println(greetingsWithName)

	var ch rune = 'a'
	jch := 'क'

	fmt.Println(ch)
	fmt.Println(jch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Types of cstr : %T\n", cstr)

	const love = "कहा चले प्यरे "
	fmt.Println(love)

}
