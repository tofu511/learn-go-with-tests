package structs_methods_interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	String() string
} 

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64  {
	return r.Width  * r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: %.2f, %.2f", r.Height, r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: %.2f", c.Radius)
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle: %.2f, %.2f", t.Base, t.Height)
}

func Perimeter(rectangle Rectangle) float64  {
	return 2 * (rectangle.Width+ rectangle.Height)
}
