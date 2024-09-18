// Package main provides tests for the Sum function, which calculates the total
// of an array of integers.
package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

// TestSum verifies that the Sum function returns the expected total for a given
// array of integers.
func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} // Define an array of integers.
		got := Sum(numbers)             // Call the Sum function with the array.
		want := 15                      // The expected result of the sum.

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers) // Error format.
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) { // You should use slices.Equal method instead.
		t.Errorf("got %v want %v", got, want) // %v prints in the default format
	}
}

func TestSumAllTails(t *testing.T) {
	assertResult := func(t testing.TB, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertResult(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertResult(t, got, want)
	})
}
