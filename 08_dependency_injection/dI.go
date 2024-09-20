package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet writes a greeting message to the provided io.Writer.
// It takes a name as a string and formats the message accordingly.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// GreeterHandler is an HTTP handler that responds with a greeting.
// It uses the Greet function to write "Hello, World" to the response.
func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	// The Greet function is designed to accept any type that implements
	// the io.Writer interface, allowing for flexibility in output.
	// This means we can use it with both bytes.Buffer for testing
	// and os.Stdout for main application output.
	//
	// Greet(os.Stdout, "Elodie") demonstrates this flexibility.
	//
	// Additionally, the http.ResponseWriter used in GreeterHandler
	// also implements io.Writer, making it compatible with Greet.
	log.Fatal(http.ListenAndServe("127.0.0.1:5001", http.HandlerFunc(GreeterHandler)))
}
