// The main package serves as the entry point for the application,
// encapsulating the functionality related to summing slices of integers.
package main

// Sum calculates the total of a slice of integers.
// It iterates through each integer in the slice, accumulating the sum.
// The function returns the total sum as an integer.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll calculates the sums of multiple slices of integers.
// It accepts a variadic number of integer slices and returns a slice
// containing the sum of each provided slice.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the sums of the tails of multiple slices of integers.
// It accepts a variadic number of integer slices and returns a slice
// containing the sum of the tail (all elements except the first) of each
// provided slice. If a slice is empty, it appends 0 to the result.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numbers[1:] // Extracts the tail of the slice.
		sums = append(sums, Sum(tail))
	}

	return sums
}

// Sidenote: a variadic function can take an arbitrary number of arguments,
// allowing for flexible input. In this case, it can accept multiple slices
// of integers, making it easier to sum different sets of numbers in one
// function call.
