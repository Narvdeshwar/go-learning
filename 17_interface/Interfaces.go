package main

import "fmt"

type Shape interface {
	Area() float32
}
type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return 2 * 3.14 * c.radius
}

type Rectangle struct {
	length, breadth float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.breadth
}

func printArea(s Shape) {
	fmt.Println("Area : ", s.Area())
}

func main() {
	newCircle := Circle{23}
	printArea(newCircle)
	newRectangle := Rectangle{2, 4}
	printArea(newRectangle)
}
