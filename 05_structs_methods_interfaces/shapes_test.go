// The main package contains the entry point for the application and is where
// the testing functions for the shapes are defined.
package main

import "testing"

// TestPerimeter verifies the perimeter calculation for a rectangle.
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // %.2f for two decimal float64
	}
}

// TestArea verifies the area calculation for different shapes using table-driven tests.
func TestArea(t *testing.T) {
	// Define test cases for area calculation with an anonymous struct
	areaTests := []struct {
		name    string // implement name for readability
		shape   Shape
		hasArea float64 // rename want to hasArea for readability
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.16},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	// Execute table-driven tests for each shape.
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// %#v format string will print our struct with its field values
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
