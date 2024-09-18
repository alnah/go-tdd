// Package iteration provides functions for string operations, including repetition.
package iteration

// Repeat returns a new string consisting of the given character repeated
// the specified number of times. If repetition is less than or equal to
// zero, an empty string is returned.
func Repeat(character string, repetition int) string {
	var repeated string
	for i := 0; i < repetition; i++ {
		repeated += character
	}
	return repeated
}
