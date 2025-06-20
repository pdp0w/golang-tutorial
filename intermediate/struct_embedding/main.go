package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person // embedded struct
	empId  string
	salary float64
}

func main() {

	emp := employee{
		person: person{
			name: "pdp0w",
			age:  34,
		},

		empId:  "E001",
		salary: 50000,
	}

	fmt.Println("Name:", emp.name) // accessing the embedded struct field emp.person.name directly
	fmt.Println("Age:", emp.age)   // same as above
	fmt.Println("empId:", emp.empId)
	fmt.Println("salary:", emp.salary)

	emp.introduce()
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID : %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
}

/*

Best Practices and Considerations
- Composition over inheritance
- Avoid diamond problem
- Clarity and Readability
- Initialization

*/
