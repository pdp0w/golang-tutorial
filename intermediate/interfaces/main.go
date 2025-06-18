package main

import (
	"fmt"
	"math"
)

// it is a contract which says, if you want this type then those type must have these methods under the interface
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{length: 4, width: 3}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
