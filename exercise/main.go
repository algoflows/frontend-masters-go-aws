package main

import "fmt"

const (
	PI = 3.14
)

type circle struct {
	radius float64
}

func createCircle(radius float64) *circle {
	return &circle{
		radius: radius,
	}
}

func (c *circle) calculateCircumference() float64 {
	return 2 * PI * c.radius
}

func (c *circle) calculateArea() float64 {
	return PI * c.radius * c.radius
}

func main() {
	newCicle := createCircle(20)
	circumfrence := newCicle.calculateCircumference()
	area := newCicle.calculateArea()
	fmt.Println("My circle's cirumfrence", circumfrence)
	fmt.Println("My cicle's area", area)
}
