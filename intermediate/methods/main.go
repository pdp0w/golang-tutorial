package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value reciver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer reciver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()

	fmt.Println(area)

	rect.Scale(2)
	area = rect.Area()

	fmt.Println(area)

	num := MyInt(-5)
	num1 := MyInt(5)

	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num1.welcomeMessage())

}

type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

func (m MyInt) welcomeMessage() string {
	return "welcome to myint type"
}
