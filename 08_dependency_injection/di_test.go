package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// The Buffer type from the bytes package implements the Writer interface,
	// because it has the method Write(p []byte) (n int, err error).
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alexis")

	got := buffer.String()
	want := "Hello, Alexis"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
