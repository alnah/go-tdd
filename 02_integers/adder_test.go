// Package integers provides functions for integer operations.
package integers

import (
	"fmt"
	"testing"
)

// TestAdder tests the Add function to ensure it returns the correct sum.
func TestAdder(t *testing.T) {
	sum := Add(2, 2) // Calculate the sum of 2 and 2
	expected := 4    // The expected result

	// Check if the sum is equal to the expected result
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum) // %d for digits
	}
}

// Adding this code will cause the example to appear in your godoc documentation.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Adding this comment means the example will also be executed.
	// Output: 6
}

// If you publish your code with examples to a public URL, 
// you can share the documentation of your code at pkg.go.dev.