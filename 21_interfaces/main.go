package main

import "fmt"

// here is a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perimeter() float64
}

// now we can define some structs where we have some properties
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// now we can define the functions for rectangles and circles
func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	// interfaces are named collection of method signatures.

	tempRectangle := rect{3, 4} // you can also write as  r := rect{width: 3, height: 4}
	tempCircle := circle{5}

	fmt.Println(tempRectangle.area())
	fmt.Println(tempCircle.area())

	fmt.Println(tempRectangle.perimeter())
	fmt.Println(tempCircle.perimeter())
}
