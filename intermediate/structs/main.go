package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func main() {

	p := Person{
		firstName: "Priyanshu",
		lastName:  "Gupta",
		age:       21,
		address: Address{
			city:    "Asansol",
			country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "984983",
			cell: "3485384593",
		},
	}

	fmt.Println(p.firstName, p.lastName, p.age)
	fmt.Println(p.address)

	// example of anonymous struct
	user := struct {
		username string
		email    string
	}{
		username: "pdp0w",
		email:    "pdp0w@pdp0w.com",
	}

	fmt.Println(user)

	// result := p.fullName()
	p.incrementAge()
	fmt.Println(p.age)
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge() {
	p.age++
}
