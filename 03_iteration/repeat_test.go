// Package iteration provides functions for string operations, including repetition.
package iteration

import (
	"fmt"
	"testing"
)

// TestRepeat verifies that the Repeat function returns the expected repeated string.
func TestRepeat(t *testing.T) {
	t.Run("repeats a character five times", func(t *testing.T) {
		repeated := Repeat("a", 5) // Call the Repeat function with the input "a".
		expected := "aaaaa"        // Define the expected output.
		assertCorrectRepetition(t, repeated, expected)
	})

	t.Run("repeat a character ten times", func(t *testing.T) {
		repeated := Repeat("l", 10)
		expected := "llllllllll"
		assertCorrectRepetition(t, repeated, expected)
	})
}

// assertCorrectRepetition checks if the repeated string matches the expected
// output. It is a helper function that simplifies error reporting in tests.
func assertCorrectRepetition(t testing.TB, repeated, expected string) {
	t.Helper() // Marks this function as a helper for better error reporting.
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated) // Error format.
	}
}

// BenchmarkRepeat measures the performance of the Repeat function by executing
// it b.N times, allowing for performance analysis and optimization.
func BenchmarkRepeat(b *testing.B) {
	// When the benchmark code is executed,
	// it runs b.N times and measures how long it takes.
	//
	// The amount of times the code is run shouldn't matter to you,
	// the framework will determine what is a "good" value.
	//
	// go test -bench=.
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

// ExampleRepeat demonstrates how to use the Repeat function to create a string
// that repeats a given character a specified number of times. This example
// is useful for understanding the function's behavior in practice.
func ExampleRepeat() {
	string := Repeat("@", 5)
	fmt.Print(string)
	//Output: "@@@@@"
}
