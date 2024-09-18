// Package main provides tests for the Hello function.
package main

import "testing" // To use the *testing.T type

// TestHello verifies that the Hello function returns the expected greeting.
func TestHello(t *testing.T) { // The test function takes that t hook arg only.
	t.Run("saying hello to people", func(t *testing.T) { // This is a subset.
		got := Hello("alnah", "") // It declares variable.
		want := "Hello, alnah"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Piccolo", "French")
		want := "Bonjour, Piccolo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João"
		assertCorrectMessage(t, got, want)
	})
}

// assertCorrectMessage is designed to be a helper for both tests and benchmarks.
// By accepting a testing.TB interface, it allows for greater flexibility.
func assertCorrectMessage(t testing.TB, got, want string) { // It shortens types.
	t.Helper() // Needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our
	// function call rather than inside our test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want) // Error format with %q placeholders.
	}
}
