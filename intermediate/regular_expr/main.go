package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am great\"")

	// compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z{2,}]`)

	// test some string
	email1 := "user@gmail.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email 1 :", re.MatchString(email1))
	fmt.Println("Email 2 :", re.MatchString(email2))

	// capturing groups
	// compiler regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// test string
	date1 := "2024-07-30"

	// find all submatches
	submatches := re.FindStringSubmatch(date1)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Target string
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	// test string
	text := "Golang is going great"

	// match
	fmt.Println("Match:", re.MatchString(text))
}
