package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	// `select`` allows you to wait on multiple channels.
	// The first one to send a value "wins".
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// We don't care what type is sent to the channel.
	// We just want to signal we are done and closing the channel.
	// `chan struct {}`` is the smallest data type available for memory.
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}