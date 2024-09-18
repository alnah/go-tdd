// Package main implements a simple program that prints a greeting message.
package main

import "fmt"

const (
	// Language constants for greeting messages.
	english    = "English"    // English language
	spanish    = "Spanish"    // Spanish language
	french     = "French"     // French language
	portuguese = "Portuguese" // Portuguese language

	// Greeting prefixes for each language.
	englishHelloPrefix    = "Hello, "   // Greeting in English
	frenchHelloPrefix     = "Bonjour, " // Greeting in French
	spanishHelloPrefix    = "Hola, "    // Greeting in Spanish
	portugueseHelloPrefix = "Olá, "     // Greeting in Portuguese
)

// Hello returns a greeting message as a string.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "" {
		language = english
	}
	return greetingPrefix(language) + name
}

// greetingPrefix returns the appropriate greeting prefix based on the provided
// language. If the language is not recognized, it defaults to English.
//
// The function name starts with a lowercase letter. 
// In Go, public functions start with a capital letter, 
// and private ones start with a lowercase letter.
//
// (prefix string) is a named return value.
func greetingPrefix(language string) (prefix string) { 
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// main is the entry point of the program. It prints "Hello, world" to the console.
func main() {
	fmt.Println(Hello("world", ""))
}
