// The main package defines shapes and their associated methods for area
// calculations and perimeter computation.
package main

import "math"

// Shape defines an interface for calculating the area of a shape.
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle represents a circle with a radius.
type Circle struct {
	Radius float64 // Radius of the circle
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Round(math.Pi*c.Radius*c.Radius*100) / 100
}

// Triangle represents a triangle with a base and height.
type Triangle struct {
	Base   float64 // Base length of the triangle
	Height float64 // Height of the triangle
}

// Area calculates the area of the triangle using the formula: (base * height) / 2.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter calculates the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}
